package main

type Tag string

const (
	PROOF_OF_STAKE    Tag = "proof-of-stake"
	PROOF_OF_WORK     Tag = "proof-of-work"
	LIQUID_STAKING    Tag = "liquid-staking"
	LIGHT_CLIENTS     Tag = "light-clients"
	INTEROPERABILITY  Tag = "interoperability"
	WALLETS           Tag = "wallets"
	PAYMENT_CHANNELS  Tag = "payment-channels"
	ZERO_KNOWLEDGE    Tag = "zero-knowledge"
	GROUP_SIGNATURES  Tag = "group-signatures"
	LIGHTNING_NETWORK Tag = "lightning-network"
)

var TagToColor = map[Tag]string{
	PROOF_OF_STAKE:    "#e4cbf4",
	PROOF_OF_WORK:     "#ffc0c0",
	LIQUID_STAKING:    "#e6e6fa",
	LIGHT_CLIENTS:     "#fff3b0",
	INTEROPERABILITY:  "#a5d0ec",
	WALLETS:           "#e1ddad",
	PAYMENT_CHANNELS:  "#e2e2e2",
	ZERO_KNOWLEDGE:    "#ffd9b4",
	GROUP_SIGNATURES:  "#fae2d9",
	LIGHTNING_NETWORK: "#d4e2d4",
}

var ResearchPapers []ResearchPaper = []ResearchPaper{
	{
		Handle:     "popos",
		Name:       "Proofs of Proof-of-Stake with Sublinear Complexity",
		Conference: "ACM Advances in Financial Technologies 2023",
		Authors:    "Shresth Agrawal, Joachim Neu, Ertem Nusret Tas, and Dionysis Zindros",
		Url:        "https://arxiv.org/pdf/2209.08673.pdf",
		Tags:       []Tag{PROOF_OF_STAKE, LIGHT_CLIENTS, INTEROPERABILITY},
	},
	{
		Handle:     "liquid-staking",
		Name:       "The Principal–Agent Problem in Liquid Staking",
		Conference: "Financial Cryptography and Data Security 2023, Workshop in Trusted Smart Contracts",
		Authors:    "Apostolos Tzinas and Dionysis Zindros",
		Url:        "https://eprint.iacr.org/2023/605.pdf",
		Tags:       []Tag{LIQUID_STAKING},
	},
	{
		Handle:     "nipopow",
		Name:       "Non-Interactive Proofs of Proof-of-Work",
		Conference: "Financial Cryptography and Data Security 2020",
		Authors:    "Aggelos Kiayias, Andrew Miller, and Dionysis Zindros",
		Url:        "https://eprint.iacr.org/2017/963.pdf",
		Tags:       []Tag{PROOF_OF_WORK, LIGHT_CLIENTS, INTEROPERABILITY},
	},
	{
		Handle:     "ouroboros",
		Name:       "Ouroboros: A provably secure proof-of-stake blockchain protocol",
		Conference: "Advances in Cryptology–CRYPTO 2017",
		Authors:    "Aggelos Kiayias, Alexander Russell, Bernardo David, and Roman Oliynykov",
		Url:        "https://eprint.iacr.org/2016/889.pdf",
		Tags:       []Tag{PROOF_OF_STAKE},
	},
	{
		Handle:     "hardware-wallets",
		Name:       "A formal treatment of hardware wallets",
		Conference: "Financial Cryptography and Data Security 2019",
		Authors:    "Myrto Arapinis, Andriana Gkaniatsou, Dimitris Karakostas, and Aggelos Kiayias",
		Url:        "https://eprint.iacr.org/2019/034.pdf",
		Tags:       []Tag{WALLETS},
	},
	{
		Handle:     "brick",
		Name:       "Brick: Asynchronous Incentive-Compatible Payment Channels",
		Conference: "Financial Cryptography and Data Security 2021",
		Authors:    "Zeta Avarikioti, Eleftherios Kokoris-Kogias, Roger Wattenhofer, and Dionysis Zindros",
		Url:        "https://arxiv.org/pdf/1905.11360.pdf",
		Tags:       []Tag{PAYMENT_CHANNELS},
	},
	{
		Handle:     "zk-arguments",
		Name:       "Efficient Zero-Knowledge Arguments for Arithmetic Circuits in the Discrete Log Setting",
		Conference: "Advances in Cryptology–EUROCRYPT 2016",
		Authors:    "Jonathan Bootle, Andrea Cerulli, Pyrros Chaidos, Jens Groth, and Christophe Petit",
		Url:        "https://eprint.iacr.org/2016/263.pdf",
		Tags:       []Tag{ZERO_KNOWLEDGE},
	},
	{
		Handle:     "dynamic-group-signatures",
		Name:       "Foundations of Fully Dynamic Group Signatures",
		Conference: "International Conference on Applied Cryptography and Network Security 2016",
		Authors:    "Jonathan Bootle, Andrea Cerulli, Pyrros Chaidos, Essam Ghadafi, and Jens Groth",
		Url:        "https://eprint.iacr.org/2016/368.pdf",
		Tags:       []Tag{GROUP_SIGNATURES},
	},
	{
		Handle:     "lightning-network-security",
		Name:       "A Composable Security Treatment of the Lightning Network",
		Conference: "IEEE Computer Security Foundations Symposium 2020",
		Authors:    "Aggelos Kiayias and Orfeas Stefanos Thyfronitis Litos",
		Url:        "https://eprint.iacr.org/2019/778.pdf",
		Tags:       []Tag{LIGHTNING_NETWORK},
	},
}
