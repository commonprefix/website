---
title: "$4M SquidRouterModule Exploit: What Exactly Happened"
date: "26/05/2026"
desc: "Analysis of a third-party Safe module exploit on May 25, 2026."
authors: "nikolas,ristic"
index: true
---

**All Axelar user funds, approvals, and integrations are unaffected, and no action is needed. Our integration partner Squid was also not affected.**

## TL;DR

- In May 2026, an exploit targeted a third-party Safe module deployed at the same address on [Ethereum](https://etherscan.io/address/0x1f1d37a3Bf840e35c6a860c7C2dA71Fe555123ca), [Base](https://basescan.org/address/0x1f1d37a3Bf840e35c6a860c7C2dA71Fe555123ca), and [Arbitrum](https://arbiscan.io/address/0x1f1d37a3Bf840e35c6a860c7C2dA71Fe555123ca): `0x1f1d37a3Bf840e35c6a860c7C2dA71Fe555123ca`.
- The observed victims were individual Safe smart accounts that had enabled the affected Safe module and granted delegate permissions through the module's permissions manager. The affected component was the affected Safe application module, not Axelar or Squid.
- Across the indexed attack transactions, the observed drain total is approximately $3.98M across 313 attack transactions, including 308 successful drains and 5 reverted attempts, and 88 distinct victim Safes, using the investigation snapshot from 26th of May 2026.
- The core issue was that the affected Safe module decoded a `delegate` address from attacker-controlled cross-chain payload bytes and then used that address for local Safe permission checks.
- After the drains, a large portion of the proceeds was consolidated across Ethereum, Base, and Arbitrum via Relay and sent as 3,070,767.09 DAI to attacker wallet 5.

## Incident Summary

On May 25, 2026, 88 distinct victim Safes were observed drained, totaling approximately $3.98M worth of stolen tokens. The exploit targeted a third-party Safe module. It was written, deployed, and operated by a third-party, and was part of a larger Safe module suite whose deployer also deployed sibling modules for Aave, Morpho, Yearn, Uniswap V2/V3/V4, and Krystal.

The identity of the module deployer has not been confirmed, but given a recent [Blockscan Chat message](https://etherscan.io/tx/0xa7db9793d4978a26ee01c2a6cbf5ba154925a7532e31ae731f7f08ba4295ee70) it's possible [NewMarketTrading](https://newmarkettrading.com) is the party behind these contracts.

The observed victims were individual Safe smart accounts. Each Safe had previously enabled the exploited module.

<br />

## How the Attack Played Out

<br />

### Step 1. The attacker identified enabled Safes and usable delegates

The affected Safe module relied on a permission model where a delegate could be authorized to perform actions for a specific Safe through a specific module. A wildcard `"*"` grant meant that a delegate could use all supported actions for that module on that Safe.

The attacker did not need to compromise the Safe owner's key or the delegate's key. Instead, the attacker selected delegates that already had broad permissions for victim Safes and then placed those delegate addresses inside attacker-controlled payloads.

### Step 2. The vulnerable handler trusted a delegate from payload bytes

The vulnerable destination-chain handler decoded the `delegate` value from the cross-chain payload:

```solidity
(address module, address safe, address delegate, ActionsExecutionParams memory params)
    = abi.decode(payload, (address, address, address, ActionsExecutionParams));

require(module == address(this), ...);
bridgedToken.safeTransfer(safe, bridgedTokenAmount);
_handleActions(safe, delegate, params);
```

The `module == address(this)` check only verified that the payload claimed to be intended for the affected Safe module. It's later checked that the delegate value had the right permissions:

```solidity
function _handleUniV3SwapExactIn(
    address safe,
    address delegate,
    bytes memory encodedData
) internal returns (bool result) {
    _checkPermission(safe, delegate, Permissions.SWAP_PERMISSION);

    // do the swap
}
```

The core issue was therefore simple: the attacker could choose the delegate address inside the payload. If that delegate already had `"*"` permission for the victim Safe, the affected Safe module accepted the request and executed the requested actions.

The observed attacks used Axelar Express Execution. That was how the calls reached the affected module, not the root cause. The vulnerable trust decision was inside the affected Safe module's payload handling: the destination contract treated attacker-controlled message fields as authority for a local Safe operation.

Some attacks called the affected Safe module directly. Others went through attacker contracts that simplified the attack call. For example, the Ethereum attacker contract [`0xfac7459683cdb9b6f367b42eedfebd745dc8760c`](https://etherscan.io/address/0xfac7459683cdb9b6f367b42eedfebd745dc8760c) decompiled to a two-argument entry point taking a victim Safe and delegate. In this write-up we refer to that behavior as `drain(safe, delegate)`, but the name is not a verified ABI name.

```solidity
// label for illustration, actual selector is 0xf8c8f0a3
drain(address safe, address delegate)
```

The attacker contract assembled the malicious `ActionsExecutionParams` internally and sent it to the affected Safe module.

The structure contained an ordered list of actions and a strictness flag. Each action carried an action type and ABI-encoded parameters interpreted by the affected Safe module:

```solidity
struct ActionsExecutionParams {
    ExecuteAction[] actions;
    bool isStrict;
}

struct ExecuteAction {
    ExecuteActionType actionType;
    bytes encodedData;
}

enum ExecuteActionType {
    UNI_V2_SWAP_EXACT_IN,
    UNI_V2_SWAP_EXACT_OUT,
    UNI_V3_SWAP_EXACT_IN,
    UNI_V3_SWAP_EXACT_OUT,
    ERC20_APPROVE,
    PERMIT2_APPROVE,
    NATIVE_WRAP,
    NATIVE_UNWRAP
}
```

Observed attacker wallets include:

- Attacker wallet 1: [`0x7c82cb4b2909c50c7c0f2b696eee7565e0a23bb8`](https://etherscan.io/address/0x7c82cb4b2909c50c7c0f2b696eee7565e0a23bb8), the main operator wallet. It deployed attacker contracts, executed most attacker-contract calls, used Relay during consolidation, and sent DAI to attacker wallet 5. It also deployed and used the fake `"u"` token (`0xe6ff0fe017d09d690493dec0f0f55e8f9cdc3512`) as the worthless token in its swap path.
- Attacker wallet 2: [`0x9bdc730183821b6bb2b51be30b77c964fa645b91`](https://etherscan.io/address/0x9bdc730183821b6bb2b51be30b77c964fa645b91), a co-operator wallet. It deployed attacker contracts, received two 1 ETH Tornado Cash withdrawals, and then funded attacker wallet 1 with operational gas across Ethereum, Base, and Arbitrum.
- Attacker wallet 3: [`0x7e54c729148a95bca651f3214ac9ebefd3fb1271`](https://etherscan.io/address/0x7e54c729148a95bca651f3214ac9ebefd3fb1271), which made direct calls to the affected Safe module. It was funded by `0x51dbd97fa14f4d78f8afc2e92a692a9e087efb5c` with 0.0687 ETH, and deployed 19 fake-token contracts across Ethereum, Base, and Arbitrum, including fake `"USDT"` tokens, together with matching Uniswap V3 pools.
- Attacker wallet 4: [`0xc8ef4003d9db3863b9af26afcf2275378bfa83e4`](https://etherscan.io/address/0xc8ef4003d9db3863b9af26afcf2275378bfa83e4), a secondary operator wallet that made direct calls to the affected Safe module. It was funded through Orbiter Finance Bridge 3, including 0.9865 ETH before the exploit window and 1.9945 ETH during the operation, and used separate fake tokens, specifically `"sasx"`, `"sas"` and `"ddxx v3"`.
- Attacker wallet 5: [`0xa447f71782135ab96a71374271a749ff7aa54859`](https://etherscan.io/address/0xa447f71782135ab96a71374271a749ff7aa54859), which received the consolidated 3,070,767.09 DAI from attacker wallet 1.

Observed attacker contracts include:

- Ethereum: [`0xfac7459683cdb9b6f367b42eedfebd745dc8760c`](https://etherscan.io/address/0xfac7459683cdb9b6f367b42eedfebd745dc8760c), deployed by attacker wallet 1, with 33 observed calls.
- Ethereum: [`0xe1d5fcfbba4d46f4937de369de415dd7e2d3265a`](https://etherscan.io/address/0xe1d5fcfbba4d46f4937de369de415dd7e2d3265a), deployed by attacker wallet 2, with 90 observed calls.
- Base: [`0x2d450322e3526f489afcc8c49923b35d355c70bc`](https://basescan.org/address/0x2d450322e3526f489afcc8c49923b35d355c70bc), deployed by attacker wallet 2, with 44 observed calls.
- Base: [`0xcc5f66da5dc7f8a70fdfcd9da610fc7e753c4867`](https://basescan.org/address/0xcc5f66da5dc7f8a70fdfcd9da610fc7e753c4867), deployed by attacker wallet 1, with 24 observed calls.
- Arbitrum: [`0x2d450322e3526f489afcc8c49923b35d355c70bc`](https://arbiscan.io/address/0x2d450322e3526f489afcc8c49923b35d355c70bc), deployed by attacker wallet 2 at the same address as Base, with 44 observed calls.
- Arbitrum: [`0x35b94f725bfdc1d00ef15df99dafadf6ecf3a0e9`](https://arbiscan.io/address/0x35b94f725bfdc1d00ef15df99dafadf6ecf3a0e9), deployed by attacker wallet 1, with 4 observed calls.

### Step 3. The Safe approved and swapped assets into attacker-controlled destinations

Once the affected Safe module blindly accepted the payload, it called the victim Safe through `execTransactionFromModule`. The Safe then approved and swapped assets through attacker-controlled routing.

The affected Safe module did not expose a generic transfer action or arbitrary Safe call. Its action menu was limited to swaps, approvals, and native wrap/unwrap operations, and approval spenders were constrained to Squid Router, supported Universal Routers, or Permit2. The attacker therefore could not simply approve an attacker wallet or transfer tokens directly out of the Safe.

The extraction path was the swap itself. By routing through attacker-created Uniswap V3 pools paired against worthless attacker-issued tokens, the Safe sent real assets into the pool and received those worthless tokens back. The attacker, as the pool LP, could then withdraw the real assets from the pool.

The attacker used multiple fake tokens and created Uniswap V3 pools as drain destinations. One example swap on Ethereum is [`0x8de614bdb7acf5dcbdfe5ce8ed17ec2a2058e7708e6a4cf44f0e523c72df24c3`](https://etherscan.io/tx/0x8de614bdb7acf5dcbdfe5ce8ed17ec2a2058e7708e6a4cf44f0e523c72df24c3):

- The victim Safe sent 32,000 apyUSD.
- The Safe received a fake `"USDT"` token: `0xd7ef4cb5c57a4362a48abf4d33c5b9470d6681e1`.
- The fake token and the corresponding Uniswap V3 pool were both deployed by attacker wallet 3.

## Who Was Affected

The affected accounts were Safe smart accounts that had enabled the affected Safe module and had matching delegate permissions. The observed victims were individual Safe smart accounts, not a single protocol treasury.

Axelar and Squid were not affected. Their users, funds, approvals, and integrations are unaffected, and no action is needed.

The observed drain total was approximately $3.98M across 313 attack transactions, including 308 successful drains and 5 reverted attempts, and 88 distinct victim Safes:

- Ethereum: 125 successful drain transactions, approximately $1,083,764 drained.
- Base: 132 successful drain transactions, approximately $2,814,305 drained.
- Arbitrum: 51 successful drain transactions, approximately $80,765 drained.
- Total: 308 successful drain transactions, approximately $3,978,833 drained.

By token, the observed stolen assets were approximately:

- USDC: about 1.18M tokens ($1.18M).
- cbBTC: about 12.4354 tokens ($961,682).
- WETH: about 181.3 tokens ($383,306).
- apyUSD: about 371,020 tokens ($371,020).
- gtUSDCp: about 257,606 tokens ($257,606).
- cbXRP: about 97,436 tokens ($132,513).
- sUSDat: about 111,098 tokens ($111,098).
- steakUSDC: about 107,347 tokens ($107,347).
- WBTC: about 1.243 tokens ($95,975).
- wTAO: about 286.5 tokens ($82,323).
- steakETH: about 25.1997 tokens ($55,951).
- SYRUP: about 45,215 tokens ($52,450).
- wstETH: about 19.3152 tokens ($50,574).
- BNKR: about 375,795 tokens ($37,580).
- sparkUSDC: about 14,579 tokens ($14,579).
- AAVE: about 151 tokens ($13,042).
- LIT: about 9,499.4 tokens ($12,349).
- ONDO: about 23,011 tokens ($9,870).
- UNI: about 2,603.6 tokens ($8,748).
- CFG: about 25,463 tokens ($7,749).
- MORPHO: about 3,009.3 tokens ($7,192).
- ENA: about 56,582 tokens ($5,718).
- FLUID: about 1,406.7 tokens ($4,220).
- USDT: about 3,533.1 tokens ($3,533).
- EUL, LBTC, ZRO, LINK, WLD, Anon, AERO, EURC, SKY, SLVon, SOL, msETH, XAUt, COOKIE, AVNT, USDT0, Cake, AUKI, and SAPIEN, together totaling about $12,171.

Base was the largest loss surface in the observed data, accounting for roughly 71% of the total. Ethereum accounted for roughly 27%, and Arbitrum accounted for roughly 2%.

Across the current dataset, 76 of the 88 observed victim Safes were drained on two or more chains.

## Fund Movement

After the drains, a large portion of the proceeds was pulled together across Ethereum, Base, and Arbitrum. Attacker wallet 1 used Relay during cross-chain consolidation and DEX/settlement infrastructure to convert proceeds before sending  3,070,767.09 DAI to attacker wallet 5 ([`0xa447f71782135ab96a71374271a749ff7aa54859`](https://etherscan.io/address/0xa447f71782135ab96a71374271a749ff7aa54859)).

At the time of the investigation notes, that DAI had not moved out of attacker wallet 5.

<br/>

## FAQs

<br/>

### Were Axelar or Squid compromised?

No. The affected component was a third-party Safe application module that incorrectly used Axelar's cross-chain execution protocol. Based on a recent [Blockscan Chat message](https://etherscan.io/tx/0xa7db9793d4978a26ee01c2a6cbf5ba154925a7532e31ae731f7f08ba4295ee70) it's possible [NewMarketTrading](https://newmarkettrading.com) is the victim platform and the author of these contracts.

### Did the attacker compromise victim Safe owners?

Victim owner keys were not compromised. The exploit used the affected Safe module's permissions and an attacker-controlled payload-provided delegate.

### Did the attacker compromise delegate keys?

Delegate private keys were not compromised. The attacker placed already-authorized delegate addresses into the payload and relied on the affected Safe module to use those addresses for permission checks.
