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

	},
}



