---
title: "XRP Ledger Batch Transactions Security Audit Report"
date: "01/09/2026"
desc: "A security audit of the XRP Ledger Batch Transactions feature (XLS-0056): seven findings across batch signing, fee calculation, execution modes, and P2P relay, including two critical and high-severity signing-scheme flaws and a DoS amplification vector."
authors: "ivan,dejan,nikolas"
index: true
---
## Abstract
This report presents the findings of a security audit of the Batch Transactions feature (XLS-0056) for the XRP Ledger, conducted by Common Prefix. The audit focused on the reference C++ implementation in the `xrpld` codebase, covering batch signing, fee calculation, execution mode semantics, ledger state isolation, serialization, consensus integration, peer-to-peer relay, replay protection, and inner transaction independence. We identified seven findings across these areas, including two critical and high-severity issues in the batch signing scheme and a high-severity DoS amplification vector via transaction ID manipulation.

![Common Prefix x XRP Ledger](/static/blog/xrpl-batch-tx-audit/xrpl_batch_tx_audit_banner.jpg)

## Overview
Common Prefix was engaged by Ripple to conduct a security audit of the Batch Transactions feature for the XRP Ledger protocol. Batch transactions enable native atomic multi-transaction bundles, allowing multiple inner transactions to be submitted and executed as a single outer batch transaction with configurable execution semantics.

The audit scope covered the following pull requests in the `rippled` C++ reference implementation: 
[PR #5060](https://github.com/XRPLF/rippled/pull/5060), 
[PR #6176](http://github.com/XRPLF/rippled/pull/6176), 
[PR #6069](http://github.com/XRPLF/rippled/pull/6069),
[PR #6446](https://github.com/XRPLF/rippled/pull/6446).

### Findings Severity Breakdown
Our findings are classified under the following severity categories, according to their impact and their likelihood of leading to an attack.

<table style="border-collapse:collapse;width:100%;font-size:0.95em;line-height:1.5">
  <thead>
    <tr style="background:#f4f1ea">
      <th style="border:1px solid #d9d4c7;padding:10px 16px;text-align:left;width:150px">Level</th>
      <th style="border:1px solid #d9d4c7;padding:10px 16px;text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px;vertical-align:top">Critical</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Logical errors or implementation bugs that are easily exploited.</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px;vertical-align:top">High</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Logical errors or implementation bugs that can be exploited but require a particular set of circumstances or can lead to DoS attacks.</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px;vertical-align:top">Medium</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Issues that may break the intended logic or deviate from specification.</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px;vertical-align:top">Low</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Issues harder to exploit, can lead to poor performance, or error-prone implementation.</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px;vertical-align:top">Informational</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Advisory comments and recommendations that could help make the codebase clearer, more readable, and easier to maintain.</td>
    </tr>
  </tbody>
</table>

---

## Findings

### CP-BATCH-01, Batch Transaction Replay Attack

**Severity: Critical**

*Description:* The batch signers' signatures are not tied to the outer account or sequence number. The signing data (constructed in `serializeBatch` in `Batch.h`) includes only the batch prefix, flags, and inner transaction IDs, but does not include the outer account, sequence number, or any unique batch identifier.

This allows an attacker who observes a valid batch transaction on the network to extract the `sfBatchSigners` array and replay the inner transactions in a new batch wrapper with different outer account credentials.

The `tfOnlyOne` execution mode is the most severely affected: it executes only the first inner transaction that succeeds, leaving the rest unexecuted. An attacker can replay the batch repeatedly, each time with conditions arranged so that a different inner transaction succeeds, effectively executing all inner transactions individually, defeating the intended mutual exclusion guarantee.

The `tfAllOrNothing` mode is also vulnerable: if a batch fails (e.g., insufficient funds), sequence numbers are not consumed. If conditions later change, the batch can be replayed since sequences remain valid.

The `tfIndependent` and `tfUntilFailure` modes are less affected because they consume sequence numbers even for failed inner transactions.

*Impact:* An attacker can replay batch transactions across different execution contexts. In `tfOnlyOne` mode, the mutual exclusion guarantee is completely broken, allowing all N inner transactions to be executed when only 1 was intended. This can lead to direct loss of funds.

```cpp
// Batch.h, serializeBatch
inline void serializeBatch(
    Serializer& msg, std::uint32_t const& flags,
    std::vector<uint256> const& txids)
{
    msg.add32(HashPrefix::batch);
    msg.add32(flags);
    msg.add32(std::uint32_t(txids.size()));
    for (auto const& txid : txids)
        msg.addBitString(txid);
    // outer account, sequence, and batch ID are NOT included
}
```

*Resolution:* `serializeBatch` now includes the outer transaction's Account and Sequence in the data signed by each batch signer, so an `sfBatchSigners` array is valid only for one specific outer submission and cannot be reused in a different batch wrapper.

---

### CP-BATCH-02, Multi-Sign Signature Reuse

**Severity: High**

*Description:* In `checkBatchMultiSign`, the batch signer's own account ID is not included in the data that individual multi-signers sign. The signing data is constructed as: batch prefix + flags + inner transaction IDs + individual `sfSigner`'s account (via `multiSignHelper`). The batch signer's account is not included.

By contrast, `checkBatchSingleSign` correctly includes the batch signer's account via `finishMultiSigningData(batchSigner.getAccountID(sfAccount), msg)`.

This means that if two batch signers (e.g., Alice and Bob) in the same batch share a common multi-signer (e.g., Carol appears on both signer lists), Carol's signature from Alice's entry can be copied into Bob's entry. The signed data is identical because the batch signer's account is not part of it.

*Impact:* A multi-signer's signature can be silently replayed across batch signer entries within the same batch transaction. This violates the fundamental property that a cryptographic signature represents explicit consent for a specific action.

```cpp
// STTx.cpp, checkBatchMultiSign
Expected<void, std::string>
STTx::checkBatchMultiSign(
    STObject const& batchSigner,
    Rules const& rules) const
{
    Serializer dataStart;
    serializeBatch(dataStart, getFlags(),
        getBatchTransactionIDs());
    return multiSignHelper(
        batchSigner,
        std::nullopt,
        [&dataStart](AccountID const& accountID)
        -> Serializer {
            Serializer s = dataStart;
            finishMultiSigningData(accountID, s); // individual multi-signer's account,
                                                  // NOT the batch signer's account
            return s;
        },
        rules);
}
```

*Resolution:* `checkBatchMultiSign` now includes the batch signer's `sfAccount` in the signing data shared by its multi-signers, matching what `checkBatchSingleSign` already did. The signed payload is now distinct per batch signer, so a multi-signer's signature can no longer be copied between batch signer entries.

---

### CP-BATCH-03, Batch Multi-Sign Self-Signing Bypass

**Severity: Low**

*Description:* In `checkBatchMultiSign`, `std::nullopt` is passed to `multiSignHelper` as the `txnAccountID` parameter. In the regular `checkMultiSign` path, the transaction's account is (conditionally) passed instead, which triggers a check in `multiSignHelper` that prevents the account owner from appearing as one of their own multi-signers.

By passing `std::nullopt`, batch multi-sign skips this restriction. A batch signer's own account can appear in its `sfSigners` array and count toward the multi-sign quorum, effectively signing for itself.

*Impact:* A batch signer who is also on their own signer list can satisfy quorum requirements with fewer independent signers than intended.

*Resolution:* Passing the batch signer's account ID instead of `std::nullopt` reactivates the check in `multiSignHelper` that prevents an account from appearing in its own `sfSigners` array. Self-signing was confirmed not to be an intentional carve-out.

---

### CP-BATCH-04, Fee Escalation Undercount for Batch Transactions

**Severity: Informational**

*Description:* During open ledger fee escalation, a batch transaction containing N inner transactions counts as only +1 toward the `txCount` that drives the escalation curve. The same N transactions submitted individually would count as +N.

The batch transaction pays a correct fee for its own admission, but exerts significantly less upward pressure on the fee curve for subsequent transactions.

With `maxBatchTxCount` = 8, a batch gets 8 transactions processed while only raising the escalation counter by 1 instead of 8.

*Impact:* During network congestion, batch transactions exert disproportionately low upward pressure on the fee escalation curve.

```cpp
// TxQ.cpp, scaleFeeLevel
auto const current = view.txCount(); // batch = +1, not +N
if (current > target)
    return mulDiv(multiplier,
        current * current,
        target * target);
```

*Resolution:* No code change. Ripple confirmed the behavior is intentional: a batch is a single transaction at the protocol layer, and the cost of its inner transactions is already charged through the aggregated fee in `Batch::calculateBaseFee` at admission time, so `txCount` deliberately remains a per-transaction count. The design rationale will be documented in the XLS-0056 specification.

---

### CP-BATCH-05, sfBatchSigners Permutation Allows Transaction ID Manipulation

**Severity: High**

*Description:* Two properties combine to create this vulnerability:

1. `sfBatchSigners` is excluded from the outer signature but included in the transaction ID. The `notSigning` flag in `sfields.macro` means the outer account's signature does not cover `sfBatchSigners`. But the transaction ID is computed via `STObject::getHash` with `withAllFields = true`, which includes `notSigning` fields.
2. `STArray` serializes elements in insertion order with no sorting applied. `STArray::add` iterates the array in order-dependent fashion. Compare with regular `sfSigners`, where `multiSignHelper` enforces ascending `AccountID` order. No equivalent ordering check exists in `checkBatchSign`.

With up to 8 batch signers (`maxBatchTxCount` = 8), this allows up to 8! = 40,320 permutations of the same valid batch transaction, each with a unique transaction ID. Each permutation has a valid outer signature (`sfBatchSigners` is `notSigning`), valid batch signer signatures (they sign flags + inner tx IDs, not signer order), and a different transaction ID (different serialization = different hash).

*Impact:* A malicious node can observe a legitimate batch transaction, generate all permutations of `sfBatchSigners`, and relay them directly to peers via the overlay. Each permutation is seen as a new transaction by the `HashRouter` and triggers preflight including cryptographic signature verification on every receiving node.

Furthermore, other permutation attacks may be exposed as Batch functionality is built on.

*Proof of Concept:* We confirmed the vulnerability on a test network with three honest validator nodes and one malicious node, all peered on a private network with the `BatchV1_1` amendment enabled. The malicious node was modified to intercept batch transactions in `NetworkOPsImp::processTransaction`, generate all permutations using `std::next_permutation`, and relay them directly via the overlay, bypassing local processing.

One batch transaction with 8 batch signers was submitted to the network:

table style="border-collapse:collapse;width:100%;font-size:0.95em;line-height:1.5">
  <thead>
    <tr style="background:#f4f1ea">
      <th style="border:1px solid #d9d4c7;padding:10px 16px;text-align:left;width:60%">Metric</th>
      <th style="border:1px solid #d9d4c7;padding:10px 16px;text-align:left">Value</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Batch signers</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">8</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Permutations generated and relayed</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">40,319</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Fee paid (victim)</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">180 drops (once)</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Attacker cost</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">zero</td>
    </tr>
    <tr>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">Batch txs in validated ledger</td>
      <td style="border:1px solid #d9d4c7;padding:10px 16px">1</td>
    </tr>
  </tbody>
</table>

Each preflight verification involves 8 cryptographic signature operations (one per batch signer). All permutations passed preflight before any sequence check. Detection is difficult because each permutation is a fully valid, properly signed transaction.

```cpp
// sfields.macro
UNTYPED_SFIELD(sfBatchSigners, ARRAY, 31,
    SField::sMD_Default, SField::notSigning)

// STObject::getHash - includes notSigning fields
uint256 STObject::getHash(HashPrefix prefix) const
{
    Serializer s;
    s.add32(prefix);
    add(s, withAllFields); // includes notSigning
    return s.getSHA512Half();
}

// No ordering check for sfBatchSigners.
// Compare: multiSignHelper enforces order for sfSigners:
//    if (lastAccountID > accountID)
//      return Unexpected("Unsorted Signers array.");
```

*Resolution:* `sfBatchSigners` is now required to be sorted by `AccountID`, matching the existing rule for `sfSigners`, via a strictly increasing `AccountID` check that also subsumes the previous duplicate check. This leaves exactly one canonical serialization, and therefore one transaction ID, per set of batch signers. Finding was accepted as a consensus-determinism issue rather than DoS amplification, which we agree.

---

### CP-BATCH-06, Missing Signature Caching for Batch Signers

**Severity: Low**

*Description:* The outer account signature of a batch transaction is cached via the `SF_SIGGOOD` flag in the `HashRouter` (`checkValidity` in `apply.cpp`). Once cached, subsequent preflight passes skip the outer signature check entirely. However, batch signer signatures are verified through a separate path that has no caching mechanism. They are re-verified from scratch on every preflight and preclaim invocation.

A batch transaction's lifecycle involves four full rounds of batch signer signature verification:

- **Stage 1, transaction submission preflight:** When a batch transaction is first received (via `PeerImp::checkTransaction` or direct submission), the outer signature is checked and cached via `SF_SIGGOOD`. The transaction then enters `tryDirectApply`, which runs preflight. `Batch::preflightSigValidated` calls `STTx::checkBatchSign`, performing the first full cryptographic verification of every batch signer's signature.
- **Stage 2, tryDirectApply preclaim:** `invoke_preclaim` in `applySteps.cpp` calls `Batch::checkSign`, which in turn calls `Batch::checkBatchSign`.
- **Stage 3, ledger close preflight:** When the transaction is applied during ledger close, preflight runs again. `Batch::preflightSigValidated` calls `STTx::checkBatchSign`.
- **Stage 4, ledger close preclaim:** `invoke_preclaim` calls `Batch::checkSign`, which calls `Batch::checkBatchSign`.

*Impact:* With 8 batch signers, each verification round performs 8 cryptographic signature operations. Across the 4 rounds, this totals 32 `verify()` calls per batch transaction. In the worst case with multi-signing (up to 32 multi-signers per batch signer), a single batch transaction triggers up to 1,024 `verify()` calls. Since the outer signature is cached via `SF_SIGGOOD` but batch signer signatures are not, this asymmetry means that batch transactions impose significantly more CPU load than their outer signature caching would suggest. During periods of high batch transaction volume, this could contribute to ledger close delays.

```cpp
// apply.cpp, checkValidity - outer signature cached
if (!any(flags & SF_SIGGOOD))
{
    auto const sigVerify = tx.checkSign(rules);
    router.setFlags(id, SF_SIGGOOD); // cached
}

// Batch.cpp, preflightSigValidated - NOT cached
auto const sigResult =
    ctx.tx.checkBatchSign(ctx.rules);

// Batch.cpp, checkSign (preclaim) - verified AGAIN
if (ctx.tx.isFieldPresent(sfBatchSigners))
{
    if (auto ret = checkBatchSign(ctx);
        !isTesSuccess(ret))
        return ret;
}
```

*Resolution:* No code change. Ripple does not consider the missing cache a security issue at current batch volumes, since the per-signer component of `Batch::calculateBaseFee` and open ledger fee escalation limit the cost of repeated submissions. Extending `SF_SIGGOOD` caching to batch signers, or adding a separate `SF_BATCHSIGGOOD` flag, is tracked as a follow-up optimization to be prioritized if profiling shows batch signer verification contributing to ledger close latency.

---

### CP-BATCH-07, Amendment Guard Gap for Inner Batch Transactions

**Severity: Low**

*Description:* In `NetworkOPs.cpp`, both `submitTransaction` and `preProcessTransaction` reject inner batch transactions only when `featureBatchV1_1` is enabled. In practice, when the amendment is not enabled, this code path is unreachable: `Transactor::preflight1` already returns `temINVALID_FLAG` for any transaction with `tfInnerBatchTxn` when the amendment is inactive. Nevertheless, the guards in `NetworkOPs` are unnecessary and inconsistent with the rest of the codebase.

By contrast, `PeerImp::handleTransaction` and `PeerImp::checkTransaction` reject `tfInnerBatchTxn` unconditionally, with an explicit comment explaining that the flag indicates an invalid transaction regardless of amendment status.

*Impact:* Before `featureBatchV1_1` activation, an attacker could submit transactions with the `tfInnerBatchTxn` flag directly to the network. While these would likely fail during processing, they would consume network resources before being rejected.

```cpp
// NetworkOPs.cpp, submitTransaction
if (iTrans->isFlag(tfInnerBatchTxn) &&
    m_ledgerMaster.getValidatedRules()
        .enabled(featureBatchV1_1))
{
    JLOG(m_journal.error())
        << "Submitted transaction invalid: "
           "tfInnerBatchTxn flag present.";
    return;
}
```

*Resolution:* Both checks in `NetworkOPs.cpp` now reject transactions carrying `tfInnerBatchTxn` regardless of amendment status, matching `Transactor::preflight1` and `PeerImp`.

---

## 3 Audit Structure

### 3.1 Batch Signing Scheme
We performed a line-by-line review of the entire batch signature verification chain.

We identified three issues in this area. The most critical is that the signing data does not include the outer batch transaction's account or sequence number, making batch signer signatures transferable across different batch wrappers (CP-BATCH-01). The second issue is that `checkBatchMultiSign` omits the batch signer's account from the data that individual multi-signers sign, while `checkBatchSingleSign` correctly includes it (CP-BATCH-02). A third, lower-severity issue is that the self-signing restriction present in regular multi-sign is bypassed in batch multi-sign (CP-BATCH-03).

### 3.2 Fee Calculation
We audited `Batch::calculateBaseFee` in `Batch.cpp`, which computes the required fee for a batch transaction. The function correctly aggregates: the default Transactor base fee, an additional `view.fees().base` for batch overhead, the sum of each inner transaction's individual base fee, and a per-signer fee based on the number of signature verification operations.

Overflow checks are present throughout the calculation. The fee is always claimed by the outer batch transaction regardless of whether inner transactions succeed.

We identified one concern with the interaction between batch transactions and the open ledger fee escalation mechanism (CP-BATCH-04).

### 3.3 Execution Mode Semantics
We traced the execution of all four batch modes through `applyBatchTransactions` in `apply.cpp`. Each mode was analyzed for atomicity, state rollback, and partial application correctness.

**tfAllOrNothing:** If any inner transaction returns a non-`tesSUCCESS` result, the function returns false immediately. The entire `wholeBatchView` is never applied, so all accumulated state changes (including sequence increments for previously-successful inner transactions) are rolled back. This correctly implements all-or-nothing atomicity.

**tfOnlyOne:** The loop breaks on the first `tesSUCCESS` inner transaction. Inner transactions that result in `tecCLAIM` still have their views applied, consuming their sequence and fee, and the loop continues past them.

**tfUntilFailure:** The loop breaks on the first non-`tesSUCCESS` result. All previously executed transactions (`tesSUCCESS` and `tecCLAIM`) have their state changes persisted. The failed transaction's changes are also applied if the result is `tecCLAIM` (consuming sequence and fee), but discarded if it is a non-tec/non-tes failure (`tem`, `tel`, `tef`).

**tfIndependent:** All inner transactions execute regardless of individual results. Both `tesSUCCESS` and `tecCLAIM` transactions have their views applied and state changes persisted. Only non-tec/non-tes failures have their views discarded.

Across all modes, each inner transaction's `perTxBatchView` is applied to the parent `batchView` whenever the result is `tesSUCCESS` or `tecCLAIM`. This means tec-class results ("claimed" transactions) consume their sequence and fee, and their state changes are persisted, just like successful transactions. Only non-tec/non-tes failures (`tem`, `tel`, `tef`) have their views discarded.

No findings were identified in the execution mode logic itself. The replay concerns affecting `tfOnlyOne` and `tfAllOrNothing` modes stem from the signing scheme (see CP-BATCH-01 in Section 2), not from the mode implementation.

### 3.4 Ledger State Isolation
We examined the ledger state isolation mechanism for inner transactions. Each inner transaction executes in its own OpenView (`perTxBatchView`), constructed with the `batch_view` tag from the shared `batchView`. This creates a two-tier isolation model.

**Tier 1, per-transaction isolation:** Each inner transaction's state modifications (account balances, sequence increments, trust line changes, offer book mutations) are confined to its `perTxBatchView`. Only on successful application are changes promoted to the parent `batchView`.

**Tier 2, whole-batch isolation:** The `batchView` accumulates all successfully-applied inner transaction changes. This view is applied to the main ledger view only if `applyBatchTransactions` returns true. For `tfAllOrNothing` mode, returning false discards the entire batch's accumulated state.

Inner transactions can see the effects of prior inner transactions in the same batch (because successful `perTxBatchViews` are applied to the shared `batchView`). This is intentional and necessary for correct sequence numbering.

No findings were identified. The two-tier OpenView model provides proper isolation with controlled promotion of state changes.

### 3.5 Serialization and Parsing
We reviewed the serialization of batch-specific fields and their interaction with transaction ID computation and signature verification.

`sfRawTransactions` (defined in `sfields.macro`) is a standard ARRAY field included in the outer transaction signature. Inner transaction data is therefore committed to by the outer account's signature, preventing modification.

`sfBatchSigners` (also in `sfields.macro`) is marked as `notSigning`, correctly excluding it from the outer transaction signature hash. However, it is included in the transaction ID computation. We identified that the lack of canonical ordering on this array creates a DoS amplification vector.

### 3.6 Consensus and Ledger Building
We traced the batch transaction's path through the consensus and ledger building pipeline. During ledger close, the batch transaction is applied via `applyTransaction` in `apply.cpp`. The outer batch transaction consumes the outer account's sequence and fee first. Inner transactions then execute within a `wholeBatchView`, and accumulated changes are applied to the main view only if `applyBatchTransactions` returns true.

The outer batch transaction is always committed to the ledger (consuming sequence and fee) regardless of inner transaction outcomes. Even if all inner transactions fail in `tfAllOrNothing` mode, the outer batch is recorded.

We checked for potential consensus divergence scenarios (different validators producing different results for the same batch) and found none. The deterministic execution model (sequential inner transaction application, deterministic view merging) ensures all validators arrive at the same result.

No findings were identified. The fee escalation undercount (CP-BATCH-04 in Section 2) has an indirect effect on ledger fill dynamics but does not threaten consensus integrity.

### 3.7 Peer-to-Peer Relay
We reviewed all entry points where batch transactions interact with the P2P relay layer. `PeerImp.cpp` provides the first line of defense: `handleTransaction` and `checkTransaction` both reject inner batch transactions received from peers, applying a moderate resource charge (`feeModerateBurdenPeer`). This is not gated behind an amendment and provides unconditional protection.

For DoS resistance, the `HashRouter` treats each batch transaction as a unique entry keyed by transaction ID. Since `sfBatchSigners` permutations produce different transaction IDs (see CP-BATCH-05 in Section 2), the `HashRouter` cannot deduplicate them.

We identified one performance concern related to signature caching (CP-BATCH-06 in Section 2).

### 3.8 Sequence, Ticket, and Replay Protection
We analyzed sequence and ticket consumption across all four execution modes.

The outer batch transaction always consumes its sequence/ticket via `Transactor::consumeSeqProxy`, called before `doApply`. This prevents the outer batch itself from being replayed.

Inner transactions consume their own accounts' sequences within isolated `perTxBatchViews`. As described in Section 3.3, both `tesSUCCESS` and `tecCLAIM` results have their views applied, meaning sequences are consumed in both cases. Sequence consumption behavior differs by mode: `tfAllOrNothing` discards the entire `wholeBatchView` on any non-`tesSUCCESS` result, rolling back all inner sequences. `tfOnlyOne` applies all views up to and including the first `tesSUCCESS` transaction (`tecCLAIM` results before it also consume their sequences). `tfUntilFailure` applies inner transactions sequentially, consuming sequences for `tesSUCCESS` and `tecCLAIM` results, but breaks on the first non-`tesSUCCESS` result (including `tecCLAIM`). A `tecCLAIM` transaction is applied (sequence consumed) and then the loop stops, so remaining transactions are never attempted. `tfIndependent` consume sequences for all `tesSUCCESS` and `tecCLAIM` inner transactions.

Batch preflight validates that inner transactions have valid sequence or ticket values, prevents duplicate sequences within the same batch, and rejects transactions with both `sfSequence` and `sfTicketSequence`.

No additional findings were identified beyond CP-BATCH-01 (Section 2), which describes how the signing scheme's lack of outer context binding undermines replay protection. The sequence and ticket validation logic itself is correct.

### 3.9 Inner Transaction Independence
We systematically verified that inner transactions cannot be extracted from a batch and submitted independently. The defense consists of multiple layers:

- **Layer 1, structural markers:** Inner transactions must have the `tfInnerBatchTxn` flag (enforced in `Batch::preflight`), zero fee, and empty `SigningPubKey`.
- **Layer 2, P2P rejection:** `PeerImp.cpp` rejects any relayed transaction with `tfInnerBatchTxn` via `handleTransaction` and `checkTransaction`, and charges a resource penalty. This is not gated behind an amendment.
- **Layer 3, network submission rejection:** `NetworkOPs.cpp` rejects inner transactions at `submitTransaction` and `preProcessTransaction`. However, both checks are gated behind `featureBatchV1_1`.
- **Layer 4, validity checking:** `checkValidity` in `apply.cpp` rejects inner batch transactions with signatures, signing public keys, or signer arrays.

### 3.10 Lending Protocol Integration
We reviewed the `LoanSet` transactor's batch handling. `LoanSet` normally requires a `sfCounterpartySignature`; for batch inner transactions, this requirement is waived and the counterparty is instead expected to authorize as a batch signer; the `sfCounterparty` field is required. Fee calculation correctly accounts for the absent counterparty signature.

No additional findings were identified.