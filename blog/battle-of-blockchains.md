---
title: "The Battle of Blockchains: PoW vs PoS Unveiled"
date: "17/04/2024"
desc: "A perspective on the long-standing battle between proof-of-work and proof-of-stake blockchains."
authors: "zeta"
draft: true
---

Since the dawn of civilization, we have managed to escape many physical
barriers of nature. With enough money, a human construct devoid of any
physical worth, one can buy the essentials to cover all their basic
needs. Yet our bodies exemplify the intangible nature of money: for
instance, the fruits of physical exercise, such as a healthier and
longer lifespan on average, can only be reaped through effort and cannot
be bought. Nonetheless, a debate arises when considering the advantages
that financial resources can bring, such as access to health supplements
or advanced exercise equipment, thus posing the question of their true
worth. *Is the essence of effort so easily substituted by financial
resources?*

The quest to answer this question is pivotal for the fabric of modern
society as well as the realm of blockchains. Indeed, it is a common
belief that blockchains operate optimally when users with more money
imply their wealth to secure blockchains. We contest this folklore
wisdom, advocating for blockchains operated via effort or ,in machine
language, computational power; the so-called Proof-of-Work (PoW). Put
simply, we argue that blockchains, much like our bodies, are more secure
with "work".

Delving into the foundations of blockchains, two properties prevail as
their backbone [@bitcoinbackboneprotocol]: safety, ensuring that
transactions confirmed by an honest node will be eventually confirmed by
all honest nodes, and liveness, guaranteeing that a transaction that is
provided to honest nodes for a long enough period of time will be
eventually confirmed by all honest nodes. These properties ensure that
blockchains produce an ordered sequence of transactions, commonly known
as a transaction ledger, which is immutable and keeps making progress.

Bitcoin made blockchains possible in a permissionless and dynamic
network where parties can join and leave at will at any point in time
without fulfilling any requirements. Newly joining nodes can trustlessly
verify the correctness of the transaction ledger, given they connect to
at least one honest node when bootstrapping in the network. In other
words, when a node joins the Bitcoin network, they may receive several
different versions of the ledger (i.e., different transaction
orderings), but they can locally identify the "correct one" as long as
it is among the received ones. Note that by "correct one" we mean the
transaction order that the honest nodes in the system have agreed upon
so far. This ability safeguards the transaction order integrity, the
manifestation of safety in a dynamic, permissionless network.

To achieve this strong property, Bitcoin leverages PoW to elect block
proposers in a fair unbiasable manner and adopts the longest chain
selection rule to determine which chain is the "correct one". PoW
effectively secures the blockchain against adversaries that control less
than 50% of the computational power of the entire system. Coupled with
the longest chain selection rule, PoW allows newly joining nodes to
identify which of the received versions of the ledger holds more work
and accept it as the "correct one".

Alternative approaches to PoW have been proposed, where the block
proposer election is proportional to the amount of coins (i.e., stake)
one owns within the system (e.g., [@kiayias2017ouroboros]). This
mechanism is known as Proof-of-Stake (PoS), and it is usually coupled
with the longest chain selection rule, akin to the one of PoW. Although
these approaches seem similar, PoS is vulnerable to safety attacks when
the network is dynamic. Let us illustrate this vulnerability with an
example. Consider a PoS blockchain that in its 100th operation day has
10% of its total stake active, with honest nodes holding 6% and an
adversary holding 4%. Then, at a later point in time, e.g. on its 200th
operation day, the total stake active is 50%, with honest nodes holding
30% and the adversary 20%. On the 200th operation day, the blockchain is
not safe anymore, as the adversary can now use its 20% stake to rewrite
*a posteriori* the transaction order of the blockchain. This is feasible
because the adversary's current controlling share is 20% of the total
stake in the system and can thus outpower the old 10% of total active
stake: adversarial block proposers corresponding to that 20% simply need
to sign fake past blocks and forge a longer chain, thus fooling newly
joining nodes to choose it as the "correct one". Remarkably, the
attacker can break the security of the blockchain swiftly and at no
cost, as no effort is required.

In essence, PoW blockchains function similarly to building a structure,
while PoS resemble the process of buying land. Given the funds, the
former demands a cumulative amount of work, whereas the latter requires
effortless signatures. As such, the safety of pure PoS systems is
predicated on the adversary controlling a minority of active stake
comparatively to any point in time, including the past.

This class of attacks, referred to as posterior attacks, are feasible
because the adversary can retroactively create past blocks. The crux of
this vulnerability is thus the absence of an *"arrow of time"*, i.e., a
temporal structure that prevents tampering with the past. This arrow of
time is inherently enforced by the very same nature of PoW but lacks in
PoS. To impose an artificial arrow of time, at a much lesser cost than
PoW, recent works propose to couple PoS with Verifiable Delay Functions
(VDFs), i.e., functions that force users to perform sequential and
non-parallelizable
operations [@boneh2018verifiable; @wesolowski2019efficient; @pietrzak2019simple].

With the use of VDFs, PoS blockchains can defend against posterior
safety attacks in dynamic settings. Still, the security of PoS protocols
employing VDFs does not match those of PoW. For example, imagine Satoshi
Nakamoto bootstrapped Bitcoin in 2009 and, at that point, they
controlled more than 50% of the power invested in Bitcoin (a very
reasonable assumption!). Today, Satoshi most probably cannot override
the Bitcoin blockchain (i.e., replace it with a longer one), as this
would entail them maintaining the majority of power on average
throughout the execution of Bitcoin (from 2009 until now). Let us now
imagine the same scenario in a PoS-VDF blockchain where the entity that
bootstraps the protocol has an initial advantage and can privately
create an alternative blockchain. In fact, PoS blockchains typically
bootstrap in a centralized manner where the "creator" controls 100% of
the circulating coins. For how long shall this entity maintain more than
50% of the stake to successfully fool newly joining nodes at a later
point in time? Alarmingly, a single moment of majority control is
sufficient to compromise safety, as the creator could initiate a private
blockchain from the genesis block, always controlling over 50% of the
stake. The integration of VDFs in PoS blockchains fails to safeguard
them against attacks where an attacker with a temporary majority creates
a private blockchain *in parallel* to the "correct" one.

Notably, the difference between these protocols is the cumulative nature
of PoW versus the non-cumulative nature of VDFs. The question that
arises is thus: can one create a blockchain that is equally secure in a
dynamic network as Bitcoin but without relying on cumulative work? We
assert that the answer to this question is no. To capture the
fundamental difference across these mechanisms, consider this thought
experiment: *Given the public key that will win the lottery for
proposing the next block, what is the required computational overhead to
actually produce the next block?*

In the context of PoW, knowing the winning public key in advance does
little to alleviate the computational workload; the system must still
undergo the same computational waste to identify the correct nonce.
Conversely, in PoS blockchains that integrate VDFs, the computational
overhead is reduced to executing the VDF with the pre-identified winning
key. In a pure PoS blockchain, this process is merely a constant
operation. We posit that this theoretical exercise constitutes a
separator among such mechanisms, encapsulating the essence of their
safety in dynamic networks; demonstrating how (proof of) work prevails
in blockchains as it does in nature.

#### Acknowledgments:

We thank Giulia Scaffino for her valuable feedback on this post.

### References

TODO
