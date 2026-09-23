package main

import "html/template"

type PodcastEpisode struct {
	Number      int
	AffiliationNumber      int
	Title       string
	Guest       string
	GuestURL       string
	Affiliation string
	AffiliationURL string
	Description template.HTML
	Resources   []ResourceLink
	Date        string // optional
	EpisodeURLs EpisodeURLs
}

type ResourceLink struct {
	Url  string
	Name string
}

type Podcast struct {
	Title           string
	Subtitle        string
	Description     template.HTML
	SpotifyUrl      string
	ApplePodcastUrl string
	YouTubeUrl      string
	LastUpdate      string
	Episodes        []PodcastEpisode
}

type EpisodeURLs struct {
	Spotify string
	ApplePodcasts string
	YouTube string
}

var HonestMajorityPodcast = Podcast{
	Title:           "Honest Majority",
	Subtitle:        "a podcast series by Common Prefix",
	Description:     template.HTML(`Honest Majority by Common Prefix explores the people and ideas shaping trust technology. Guests share their journey into the space, the technical problems they are solving, and where they see the next five years of secure, scalable, and interoperable distributed systems headed.`),
	SpotifyUrl:      "https://open.spotify.com/show/3dZtQUzhbcnv9I5qZmaRkF?si=2996b42327d444fa",
	ApplePodcastUrl: "https://podcasts.apple.com/us/podcast/honest-majority-by-common-prefix/id1860881793",
	YouTubeUrl:      "https://www.youtube.com/@CommonPrefix",
	LastUpdate:      "December 12, 2025",
	Episodes: []PodcastEpisode{
		{
			Number:      1,
			AffiliationNumber: 2,
			Title:       "Can Data Availability be Zero-Knowledge?",
			Guest:       "Alex Evans",
			GuestURL:    "https://x.com/alexhevans",
			Affiliation: "Bain Capital Crypto",
			AffiliationURL: "https://x.com/baincapcrypto",
			Description: template.HTML(`In the inaugural episode of Honest Majority, we speak with Alex Evans (partner at Bain Capital Crypto) about the evolving relationship between data availability and zero-knowledge proofs in blockchain systems. 
The conversation begins with Alex’s path into crypto and research, before diving into the fundamentals of data availability and recent advances in zero-knowledge data availability constructions. We examine the theoretical guarantees these approaches offer, the trade-offs they introduce in terms of assumptions and complexity, and where current research still falls short. The episode concludes with a discussion of open problems and how these ideas may shape the design of future blockchain protocols.`),
			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/1809.09044",
					Name: "⁠Fraud and Data Availability Proofs: Maximising Light Client Security and Scaling Blockchains with Dishonest Majorities",
				},
				{
					Url:  "https://eprint.iacr.org/2025/034.pdf",
					Name: "ZODA: Zero-Overhead Data Availability",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/432Fe1a4yZPYUbZXztzaKo?si=2a8f4dab29364550",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/alex-evans-can-data-availability-be-zero-knowledge/id1860881793?i=1000741007403",
				YouTube: "https://youtu.be/YhChZn2s9Uc?si=1otW6qokJ_vOc8Q5",
			},
		},
		{
			Number:      2,
			AffiliationNumber: 3,
			Title:       "How Does Static Analysis Secure Smart Contracts?",
			Guest:       "Yannis Smaragdakis",
			GuestURL:    "https://x.com/YSmaragdakis",
			Affiliation: "Dedaub & University of Athens",
			AffiliationURL: "https://x.com/dedaub",
			Description: template.HTML(`In the second episode of Honest Majority, we speak with Yannis Smaragdakis (co-founder of Dedaub & Professor at the University of Athens) about compilers, program analysis, and why software foundations matter for blockchain security. We discuss how compiler theory and static analysis translate into practical tooling, where today’s smart-contract tooling still falls short, and what it would take to make correctness and security guarantees part of the default developer workflow. The conversation explores the gap between academic techniques and production systems, and why closing it is critical for the next generation of blockchain infrastructure.`),
			Resources: []ResourceLink{
				{
					Url:  "https://app.dedaub.com/decompile",
					Name: "Dedaub Decompiler",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/1VV1QVLxfUGy3cNPAVNmin?si=qXNi67QGRpyEN1JVhsfU4w",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/yannis-smaragdakis-how-does-static-analysis-secure/id1860881793?i=1000742472743",
				YouTube: "https://youtu.be/Me2CASD7fRA",
			},
		},
		{
			Number:      3,
			AffiliationNumber: 4,
			Title:       "Cardano Under Attack and Provable Security?",
			Guest:       "Prof. Aggelos Kiayias",
			GuestURL:    "https://x.com/sol3gga",
			Affiliation: "University of Edinburgh & IOG",
			AffiliationURL: "https://www.iog.io/research",
			Description: template.HTML(`In the third episode of Honest Majority, we speak with Aggelos Kiayias, Professor at the University of Edinburgh and Chief Scientist at IOG,
			about the foundations of blockchain consensus and how formal security assumptions shape real-world protocols. 
			The conversation covers Aggelos’ path into cryptography, his early work on modeling Bitcoin’s consensus, and the motivation behind Proof-of-Stake systems that do not rely on hashing power. 
			Using the recent Cardano incident as a case study, the discussion then covers disagreement among honest nodes, temporary adversarial majorities, and the self-healing properties of the Ouroboros protocol.  
			The episode concludes with a forward-looking view on privacy-enhanced smart contracts, zero-knowledge techniques, and the need to design blockchain infrastructure that remains secure in a post-quantum setting.
			`),

			Resources: []ResourceLink{
				{
					Url:  "https://eprint.iacr.org/2016/889.pdf",
					Name: "Ouroboros: A Provably Secure Proof-of-Stake Blockchain Protocol",
				},
				{
					Url:  "https://www.iog.io/papers/consensus-redux-distributed-ledgers-in-the-face-of-adversarial-supremacy",
					Name: "Consesus Redux: Distributed Ledgers in the Face of Adversarial Supremacy",
				},
				{
					Url:  "https://cardanofoundation.org/blog/november-2025-cardano-shows-resilience",
					Name: "Cardano November 2025 incident blog post",
				},

			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/5pwSFbh2FMXsmeDpIcYAP8?si=55c48a4e0a6b4ab4",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/aggelos-kiayias-cardano-under-attack-and-provable/id1860881793?i=1000744450261",
				YouTube: "https://youtu.be/Sqjz77ygpyw",
			},
		},
		{
			Number:      4,
			AffiliationNumber: 5,
			Title:       "Killing Wall Street Flash Boys with Pod",
			Guest:       "Shresth Agrawal",
			GuestURL:    "https://x.com/shresth3103",
			Affiliation: "Pod Network",
			AffiliationURL: "https://x.com/poddotnetwork",
			Description: template.HTML(`
			In the fourth episode of Honest Majority, we speak with Shresth Agrawal, Co-Founder and CEO of Pod Network, about building fast, fair, and global markets from first principles.
			The conversation traces Shresth’s path into blockchain research and engineering, including early work on light clients, Common Prefix’s formative years, and the research collaboration that led to Pod’s origins. From a Flashbots grant on decentralized auctions, the discussion follows how a simple protocol evolved into a broader vision for high-performance, leaderless market infrastructure.
			We then dive into Pod’s core technical ideas: eliminating blocks and leaders, confirming transactions in a single network round trip, and shifting bottlenecks from protocol logic to operating systems, networking, and memory. Shresth explains Pod’s central limit order book, the engineering behind achieving hundreds of thousands of orders per second, and why fairness and openness must be enforced at the protocol level, not assumed by market structure.
			The episode concludes with a discussion on where blockchain infrastructure delivers real value today, why Pod frames its mission around global markets rather than Web3, and how markets and payments may drive adoption in the years ahead.
			`),

			Resources: []ResourceLink{
				{
					Url:  "https://pod.network/",
					Name: "Pod Network",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/2eQY5dXL9s3fml2vELDnyI?si=Xkc1rqFFSPO1wQ1yzV6nLw",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/shresth-agrawal-killing-wall-street-flash-boys-with-pod/id1860881793?i=1000746051995",
				YouTube: "https://youtu.be/IuxxxmeqF8I",
			},
		},
		{
			Number:      5,
			AffiliationNumber: 6,
			Title:       "The Story Behind Mysticeti, Sui’s Consensus Algorithm",
			Guest:       "Alberto Sonnino",
			GuestURL:    "https://x.com/alberto_sonnino",
			Affiliation: "Mysten Labs",
			AffiliationURL: "https://x.com/Mysten_Labs",
			Description: template.HTML(`
			In the fifth episode of Honest Majority, we speak with Alberto Sonnino, Research Scientist at Mysten Labs, about the design and evolution of Mysticeti, Sui’s consensus protocol. 
			Alberto traces his path from distributed systems research to building production blockchain infrastructure, including lessons from moving from paper prototypes to real-world systems. 
			We unpack Mysticeti’s core ideas, DAG-based consensus, latency and throughput goals, and how it differs from earlier BFT and leader-based approaches, and discuss how consensus design interacts with networking, execution, and broader system assumptions. 
			We close with Alberto’s recommended reading for getting up to speed, plus a look ahead at what matters most next: execution, privacy, UX, and building real applications.
			`),
			Resources: []ResourceLink{
				{
					Url: "https://arxiv.org/pdf/2310.14821",
					Name: "Mysticeti: Reaching the Latency Limits with Uncertified DAGs",
				},
				{
					Url:  "https://arxiv.org/pdf/2003.11506",
					Name: "FastPay: High-Performance Byzantine Fault Tolerant Settlement",
				},
				{
					Url:  "https://arxiv.org/pdf/2209.05633",
					Name: "Bullshark: The Partially Synchronous Version",
				},
				{
					Url:  "https://dl.acm.org/doi/epdf/10.1145/3492321.3519594",
					Name: "Narwhal and Tusk: A DAG-based Mempool and Efficient BFT Consensus",
				},
				{
					Url:  "https://arxiv.org/pdf/1803.05069",
					Name: "HotStuff: BFT Consensus in the Lens of Blockchain",
				},
				{
					Url:  "https://arxiv.org/pdf/2401.10369",
					Name: "Autobahn: Seamless high speed BFT",
				},
				{
					Url:  "https://eprint.iacr.org/2024/472.pdf",
					Name: "Sailfish: Towards Improving the Latency of DAG-based BFT",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/0rXN6fnmvwkZZo6wUkhlAy?si=MrlAEEXZSWyQSqOSkoqjAw",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/alberto-sonnino-the-story-behind-mysticeti-suis/id1860881793?i=1000748535332",
				YouTube: "https://youtu.be/nuVSIOw5CA8",
			},
		},
		{
			Number:      6,
			AffiliationNumber: 7,
			Title:       "Making Bitcoin Useful Again with Trustless Lending",
			Guest:       "Vitalis Salis",
			GuestURL:    "https://x.com/vitsalis",
			Affiliation: "Babylon",
			AffiliationURL: "https://x.com/babylonlabs_io",
			Description: template.HTML(`
			In the sixth episode of Honest Majority, we speak with Vitalis Salis, Head of Engineering at Babylon, about making Bitcoin useful again through trustless lending. 
			Using the example of a BTC holder taking a stablecoin loan while keeping Bitcoin as collateral, the conversation explores what “trustless” means in practice and how enforceable guarantees can be designed without introducing custody.
			We discuss Trustless Bitcoin Vaults, how BTC can be locked into UTXOs under predefined conditions, and how supported-chain outcomes can be used to coordinate lending logic. 
			The episode dives into Bitcoin’s constraints, BitVM and off-chain enforcement mechanisms, and the broader implications of extending Bitcoin’s security model beyond passive holding.
			`),

			Resources: []ResourceLink{
				{
					Url:  "Babylon - Open Positions",
					Name: "https://babylonlabs.io/jobs",
				},
				{
					Url:  "BitVM",
					Name: "https://bitvm.org/",
				},

			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/4ymQq3VlSbtee4kXrrbs45?si=w6SItojTRie_fEf9HlUcag",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/vitalis-salis-making-bitcoin-useful-again-with/id1860881793?i=1000751548840",
				YouTube: "https://youtu.be/B37I0IDOuGQ",
			},
		},
		{
			Number:      7,
			AffiliationNumber: 8,
			Title:       "Are Faster Block Times Eating up MEV?",
			Guest:       "Jason Milionis",
			GuestURL:    "https://x.com/jason_of_cs",
			Affiliation: "Category Labs",
			AffiliationURL: "https://x.com/category_xyz",
			Description: template.HTML(`
			In the seventh episode of Honest Majority, we speak with Jason Milionis, Senior Researcher at Category Labs, about the economics of decentralized exchanges and automated market makers. 
			The conversation focuses on Loss-Versus-Rebalancing (LVR), a concept that captures the hidden cost faced by liquidity providers more precisely than impermanent loss. 
			Using the mechanics of AMMs as a starting point, we explore how prices are updated without external information, how arbitrageurs extract value, and how block times influence the frequency of updates and competition. 
			The episode also touches on broader implications for market design, MEV, and the distribution of value across participants in these on-chain systems.
			`),

			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/2208.06046",
					Name: "Automated Market Making and Loss-Versus-Rebalancing",
				},
				{
					Url:  "https://arxiv.org/pdf/2305.14604",
					Name: "Automated Market Making and Arbitrage Profits in the Presence of Fees",
				},

			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/2xfgR1zDIAR5GvkQEWlCET?si=MKCvh7PsTxOr2P9AQO0UmA",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/jason-milionis-are-faster-block-times-eating-up-mev/id1860881793?i=1000755974884",
				YouTube: "https://youtu.be/Brdxonk-Pzo?si=Xnjx9NKaEnjwkoE7",
			},
		},
		{
			Number:      8,
			AffiliationNumber: 9,
			Title:       "Coding Techniques in Walrus for Massive Storage",
			Guest:       "Lefteris Kokoris-Kogias",
			GuestURL:    "https://x.com/LefKok",
			Affiliation: "Mysten Labs",
			AffiliationURL: "https://x.com/Mysten_Labs",
			Description: template.HTML(`
			In the eighth episode of Honest Majority, we speak with Lefteris Kokoris-Kogias of Mysten Labs about the coding techniques behind Walrus, a decentralized storage protocol designed for massive data. 
			Starting from the challenges of existing storage systems, the conversation explores how Walrus uses a two‑dimensional coding scheme based on Reed–Solomon codes to achieve self‑healing, efficiently reconfigurable storage without having to constantly move large volumes of data. 
			We discuss how this design enables cheap, robust storage for use cases like NFTs, websites, and large media, and how Seal, a privacy and key‑management layer built on multi‑party computation, can provide strong cryptographic guarantees without forcing users to manage keys themselves.
			`),
			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/2505.05370",
					Name: "Walrus: An Efficient Decentralized Storage Network",
				},
				{
					Url:  "https://github.com/MystenLabs/seal/blob/main/docs/content/Seal_White_Paper_v2.pdf",
					Name: "Seal: Decentralized Secrets Management",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/22Yrqi0qxqeLHdYZ5y6Akx",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/lefteris-kokoris-kogias-coding-techniques-in-walrus/id1860881793?i=1000757737489",
				YouTube: "https://youtu.be/j2AGlgfcGQ0",
			},
		},
		{
			Number:      9,
			AffiliationNumber: 10,
			Title:       "Is FasterPay faster than FastPay?",
			Guest:       "Giorgos Tsimos",
			GuestURL:    "https://x.com/giorgos_tsimos",
			Affiliation: "Pod Network",
			AffiliationURL: "https://x.com/poddotnetwork",
			Description: template.HTML(`
			In the ninth episode of Honest Majority, we speak with Giorgos Tsimos, researcher at Pod Network, about the latest research directions at Pod, the protocol's design, and how Pod tackled the challenge of recovery mechanisms in payment systems. 
			The conversation explores a new approach where validators can process transactions without waiting for each previous one to finalize, making the system faster and more efficient without sacrificing consistency. 
			We also discuss how a single equivocation — even an accidental one — can permanently freeze a user's account, and how Pod addressed this by implementing a recovery mechanism inspired by the Simplex protocol, ensuring that an accidental equivocation never permanently costs a user their account access.
			`),
			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/2003.11506",
					Name: "FastPay: High-Performance Byzantine Fault Tolerant Settlement",
				},
				{
					Url:  "https://arxiv.org/pdf/2501.14931",
					Name: "Pod: An Optimal-Latency, Censorship-Free, and Accountable Generalized Consensus Layer",
				},
				{
					Url:  "https://eprint.iacr.org/2023/463.pdf",
					Name: "Simplex Consensus: A Simple and Fast Consensus Protocol",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/2rBFghHz1UmSirasrWKhhj",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/giorgos-tsimos-is-fasterpay-faster-than-fastpay/id1860881793?i=1000764136606",
				YouTube: "https://youtu.be/ilp5PVA0Tk8",
			},
		},
		{
			Number:      10,
			AffiliationNumber: 11,
			Title:       "Can You Trust Your Smart Contracts 100%?",
			Guest:       "Zoe Paraskevopoulou",
			GuestURL:    "https://zoep.github.io/",
			Affiliation: "National Technical University of Athens",
			AffiliationURL: "https://www.ntua.gr/en/",
			Description: template.HTML(`
			In the tenth episode of Honest Majority, we speak with Zoe Paraskevopoulou, Assistant Professor at the National Technical University of Athens and former member of the Ethereum Foundation's formal verification team, about the importance of formal verification for smart contract security. 
			The conversation begins with Zoe's path from programming languages and formal methods research into the blockchain space, before diving into what formal verification actually is and why it goes beyond testing and code review. 
			We discuss how mathematical proofs can guarantee smart contract correctness for all possible inputs and states, and explore ACT v0.2.0, a formal specification and verification framework for EVM smart contracts developed at Argot Collective. 
			The episode covers how ACT models contracts as state transition systems, how it automatically proves equivalence between bytecode and formal specifications, and the open challenges ahead including unknown code and re-entrancy vulnerabilities. 
			The conversation closes with Zoe's vision for the future, where LLMs could make formal verification more accessible and software increasingly comes with machine-checked proofs of correctness.
			`),
			Resources: []ResourceLink{
				{
					Url:  "https://github.com/argotorg/act",
					Name: "Act: Smart Contract Specification Language",
				},
			},
			EpisodeURLs: EpisodeURLs{
				Spotify: "https://open.spotify.com/episode/5gSmWwnZslex0L8ut4bFXs",
				ApplePodcasts: "https://podcasts.apple.com/us/podcast/zoe-paraskevopoulou-can-you-trust-your-smart-contracts-100/id1860881793?i=1000768413022",
				YouTube: "https://youtu.be/bnSVJTKdbdM",
			},
		},
		{
			Number:      11,
			AffiliationNumber: 12,
			Title:       "Do Fees in Rollups Bring Attacks?",
			Guest:       "Stefanos Chaliasos",
			GuestURL:    "https://x.com/schaliasosvons",
			Affiliation: "ZK Security & UCL",
			AffiliationURL: "https://x.com/zksecurityXYZ",
			Description: template.HTML(`
			In the eleventh episode of Honest Majority, we speak with Stefanos Chaliasos, Senior Security Researcher at ZK Security and Assistant Professor at UCL, about how mispriced fees in major rollups can enable practical denial-of-service and finality attacks. 
			The conversation begins with Stefanos’s journey from systems security and programming languages into blockchain and zero-knowledge proofs, before diving into the different fee vectors a rollup must handle: L2 gas execution, data availability costs, and settlement fees. 
			We discuss how mispricing these fees opens up two classes of attacks: DA Saturation, which floods a rollup’s data availability, delays finality, and reduces throughput, and Prover Killers, which exploit instructions that are cheap to execute but expensive to prove, potentially stalling provers in ZK rollups. 
			The episode also covers the challenges of conducting responsible security research and the critical security risks zero-knowledge proofs face on the path to mainstream adoption. 
			The conversation closes with Stefanos’s vision for more accessible ZKP tooling, stronger security practices, and privacy-preserving blockchain systems.
			`),
			Resources: []ResourceLink{
				{
					Url:  "https://arxiv.org/pdf/2509.17126",
					Name: "Unaligned Incentives: Pricing Attacks Against Blockchain Rollups",
				},
				{
					Url:  "https://eprint.iacr.org/2024/889.pdf",
					Name: "Analyzing and Benchmarking ZK-Rollups",
				},
			},
			EpisodeURLs: EpisodeURLs{
				YouTube: "https://youtu.be/AZvCGMqt5tA",
			},
		},
		{
			Number:      12,
			AffiliationNumber: 13,
			Title:       "The New Era of Axelar",
			Guest:       "Nikolaos Kamarinakis",
			GuestURL:    "https://x.com/nikolaskama",
			Affiliation: "Common Prefix",
			AffiliationURL: "https://x.com/CommonPrefix",
			Description: template.HTML(`
			In the twelfth episode of Honest Majority, we speak with Nikolaos Kamarinakis, Co-Founder and CTO of Common Prefix, about how Common Prefix became the lead steward of Axelar. 
			The conversation looks back on the early years of Common Prefix and its long collaboration with Axelar, which began in 2021 and led to taking over the network after the acquisition of Interop Labs in December 2025. 
			Nikolas reflects on the first months of running the network in production and the work that went into stabilising it. 
			He then looks ahead to the new Axelar App, which makes the infrastructure that has always run behind the scenes something users can access directly.
			`),
			Resources: []ResourceLink{
				{
					Url:  "https://www.commonprefix.com/blog/axelar-roadmap",
					Name: "The New Era of Axelar: 2026 Roadmap",
				},
				{
					Url:  "https://app.axelar.network",
					Name: "Axelar App",
				},
			},
			EpisodeURLs: EpisodeURLs{
				YouTube: "https://youtu.be/B6fFBgiy2Ic",
			},
		},
	},
}



