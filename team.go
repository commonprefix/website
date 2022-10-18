package main

import "html/template"

var Members map[string]TeamMember = map[string]TeamMember{
	"shresth": {
		Handle:         "shresth",
		Name:           "Shresth Agrawal",
		Specialization: "smart contract development, auditing, algorithms",
		Desc: template.HTML(`
			Shresth is a smart contract and backend developer.
			He holds a bachelor degree in Computer Science from Jacobs University Bremen.
			He has experience building efficient and secure algorithms,
			protocols, and smart contracts for several DeFi protocols.
			Previously he worked at <a href="https://www.paraswap.io/">ParaSwap</a> where he was responsible for developing a significant part of the core aggregation algorithm.
			He is interested in Cryptography, Security, Consensus Protocols, Decentralised Finance, and Ethereum.
		`),
		Image: "shresth_agrawal.jpg",
	},
	"zeta": {
		Handle:         "zeta",
		Name:           "Dr. Zeta Avarikioti",
		Specialization: "distributed systems, scaling blockchains, cryptoeconomic incentive analysis",
		Desc: template.HTML(`
			Zeta is an upcoming Assistant Professor at the Technical University of Vienna (TU Wien) in Austria.
			She is currently a post-doctoral blockchain researcher at TU Wien working with professor Matteo Maffei.
			She graduated with a PhD from ETH Zürich and holds an engineering degree from the
			National Technical University of Athens and a masters degree from National Kapodistian University of Athens.
			She specializes in distributed systems, scaling blockchains via sharding and channels,
			and the analysis of cryptoeconomic incentives.
			Among other venues, she has published in Financial Crypto, AFT and SODA.
			<a href="https://arxiv.org/pdf/1905.11360.pdf">
			  BRICK: Asynchronous Payment Channels
			</a>
			<a href="https://eprint.iacr.org/2019/1092.pdf">
			  Cerberus Channels: Incentivizing Watchtowers for Bitcoin
			</a>, and
			<a href="https://arxiv.org/pdf/1910.10434.pdf">
			  Divide and Scale: Formalization of Distributed Ledger Sharding Protocols
			</a>
			.
		`),
		Image: "zeta_avarikioti.jpg",
	},
	"aleksis": {
		Handle:         "aleksis",
		Name:           "Aleksis Brezas",
		Specialization: " smart contract development, algorithmic trading, programming languages, software engineering, testing",
		Desc: template.HTML(`
			Aleksis is a senior software engineer based in Thessaloniki. With a dozen years of software engineering
			experience under his belt, he has worked with numerous programming languages and systems, with a particular
			taste for functional programming languages. He has worked on scaling web applications and in integrating
			<a href="https://geekbot.com/">geekbot</a>, and <a href="https://dailysteals.com/">dailysteals</a>.
			His blockchain development experience includes smart contract development as well as cryptocurrency algorithmic trading.
		`),
		Image: "aleksis_brezas.jpg",
	},
	"pyrros": {
		Handle:         "pyrros",
		Name:           "Dr. Pyrros Chaidos",
		Specialization: "zero-knowledge proofs, voting",
		Desc: template.HTML(`
			Pyrros is a post-doctoral blockchain researcher at the University of Athens. He holds a PhD from the
			University College of London, an MSc in Information Security from the University College of London, an MSc in
			Theoretical Computer Science from the University of Athens, and a BA in Mathematics from the University of
			Athens. He specializes in zero knowledge proof systems, with applications on proof-of-stake blockchains and
			voting in particular. Among other venues, he has published in EUROCRYPT, ACM CCS, PKC, ESORICS, and the <a href="https://eprint.iacr.org/2021/916.pdf">Mithril: Stake-based Threshold Multisignatures</a>, <a href="https://eprint.iacr.org/2016/263.pdf">Efficient Zero-Knowledge Arguments for Arithmetic Circuits in the Discrete Log Setting</a>, and <a href="https://eprint.iacr.org/2016/368.pdf">Foundations of Fully Dynamic Group Signatures</a>.
		`),
		Image: "pyrros_chaidos.jpg",
	},
	"nikolas": {
		Handle:         "nikolas",
		Name:           "Nikolaos Kamarinakis",
		Specialization: "smart contract development, testing, software engineering, offensive security",
		Desc: template.HTML(`
			Nikolas is a software engineer based in Athens, Greece and Maryland, USA.
			He is currently pursuing a bachelor's degree in Computer Science at the University of Maryland,
			specializing in Cybersecurity.
			He has multiple years of experience in full-stack software engineering and open-source development,
			as well as some expierience in offensive security.
			Nikolas is currently focused on smart contract development at Common Prefix
			as well as at the <a href="https://researchdao.io">Research DAO</a>,
			where he also coordinates operations.
		`),
		Image: "nikolaos_kamarinakis.jpg",
	},
	"dimitris": {
		Handle:         "dimitris",
		Name:           "Dr. Dimitris Karakostas",
		Specialization: "cryptocurrency wallets, macroeconomics, checkpointing",
		Desc: template.HTML(`
			Dimitris is a post-doctoral blockchain researcher at the University of Edinburgh where he works with professor
			Aggelos Kiayias. He holds a PhD from the University of Edinburgh and an Electrical and Computer Engineering
			degree from the National Technical University of Athens. He specializes in cryptocurrency wallets and
			macroeconomics, as well as chain checkpointing. Among other venues, he has published in Financial Crypto and
			<a href="https://eprint.iacr.org/2019/034.pdf">A Formal Treatment of Hardware Wallets</a>, <a href="https://arxiv.org/pdf/1907.02434.pdf">Cryptocurrency Egalitarianism: A Quantitative Approach</a>, and <a href="https://eprint.iacr.org/2020/173.pdf">Securing Proof-of-Work Ledgers via Checkpointing</a>.
		`),
		Image: "dimitris_karakostas.jpg",
	},
	"lefteris": {
		Handle:         "lefteris",
		Name:           "Prof. Dr. Lefteris Kokoris-Kogias",
		Specialization: "blockchain scalability, decentralized randomness generation, software update dispersion",
		Desc: template.HTML(`
			Lefteris is an assistant professor at IST Austria based in Vienna. He holds a PhD from EPFL, and an Electrical
			and Computer Engineering degree from the National Technical University of Athens. He specializes in the
			systems aspects of blockchains with a focus on scalability, decentralized randomness generation, and software
			update dispersion. Among other venues, he has published in IEEE S&P (Oakland, USENIX Security, ACM CCS, PODC, <a href="https://infoscience.epfl.ch/record/255586/files/OmniLedger.pdf">Omniledger: A secure, scale-out, decentralized ledger via sharding</a>, <a href="https://scholar.google.com/citations?view_op=view_citation&hl=en&user=7TOIlqkAAAAJ&citation_for_view=7TOIlqkAAAAJ:kw52XkFRtyQC">Enhancing bitcoin security and performance with strong consistency via collective signing</a>, and <a href="https://discovery.ucl.ac.uk/id/eprint/10116632/1/Jovanovic_randomness.pdf">Scalable bias-resistant distributed randomness</a>.
		`),
		Image: "lefteris_kokoris-kogias.jpg",
	},
	"stavros": {
		Handle:         "stavros",
		Name:           "Stavros Korokithakis",
		Specialization: "DevOps, OpSec, backend development",
		Desc: template.HTML(`
			Stavros is a software developer based in Thessaloniki, Greece. He holds a bachelor's degree in Computer Science
			and Electronics Engineering from the Technical University of Crete. He has been programming for twenty-five years,
			with a focus on designing, developing, scaling, and maintaining web backends as well as on systems and data security.
			Stavros has worked in various sectors such as security, developing <a href="https://www.silentcircle.com/products-and-solutions/silent-phone">Silent Phone</a>, alongside Phil Zimmermann and Jon Callas, in healthcare, and in blockchain, developing the Internet Computer at DFINITY.
		`),
		Image: "stavros_korokithakis.jpg",
	},
	"orfeas": {
		Handle:         "orfeas",
		Name:           "Dr. Orfeas Stefanos Thyfronitis Litos",
		Specialization: "software engineering, payment channels, voting, cryptoeconomic incentive analysis",
		Desc: template.HTML(`
		Orfeas obtained his
		<a
			href="https://gitlab.com/orfeasLitos/thesis/-/raw/master/thesis.pdf?inline=false"
		>PhD in Cryptography and Blockchains</a> at the University of Edinburgh in 2021. He has worked on <a href="http://fc17.ifca.ai/preproceedings/paper_37.pdf">building and analyzing</a> decentralized applications on blockchains,
		<a href="https://fc20.ifca.ai/preproceedings/132.pdf">layer-2 protocols</a>,
		<a href="https://eprint.iacr.org/2021/747.pdf">payment channels</a>, and
		<a href="https://tokenomics2019.org/Documents/proceedings-full.pdf#page=17">voting solutions</a>, all through formal cryptographic methods. Among others, he has
		<a href="https://ieeexplore.ieee.org/abstract/document/9155145">formally analyzed the security of the Lightning Network</a> and created a
		<a href="https://eprint.iacr.org/2021/747.pdf">novel virtual payment channel construction</a>.
		He is knowledgeable in
		<a
			href="https://github.com/OrfeasLitos/virtual-payment-channels/tree/master/simulation"
		>software engineering</a>
		and <a href="https://github.com/OrfeasLitos/cerberus-script">secure architecture</a>.
		His interests further include formal verification, <a href="http://fc22.ifca.ai/preproceedings/69.pdf">incentive analysis</a>, and provable security.
		`),
		Image: "orfeas_litos.jpg",
	},
	"pkakelas": {
		Handle:         "pkakelas",
		Name:           "Dimitris Lamprinos",
		Specialization: "smart contract development, large scale system design, software engineering, consensus, DevOp",
		Desc: template.HTML(`
			Dimitris is a software engineer based in Thessaloniki. He holds a Bachelor's degree in Computer Science from Aristotle
			University of Thessaloniki. He works on smart contract development and basic consensus development. He has
			significant experience in developing and scaling web applications as well as blockchain development in the
			<a href="https://www.amondo.com/">Amondo</a>, and <a href="https://geekbot.com/">Geekbot</a>.
		`),
		Image: "dimitris_lamprinos.jpg",
	},
	"themis": {
		Handle:         "themis",
		Name:           "Themis Papameletiou",
		Specialization: "smart contract development, testing, software engineering, algorithmic trading, rockets",
		Desc: template.HTML(`
			Themis is a software engineer based in Athens.
			He has significant experience developing software for a variety of projects
			such as web applications, algorithmic traders and rockets.
			He is pursuing a Master's degree in Electrical and Computer Engineering
			from the National Technical University of Athens
			and has also completed internships at Google and at the European Space Agency.
			Currently, he is focused on smart contract development and testing.
		`),
		Image: "themis_papameletiou.jpg",
	},
	"apostolos": {
		Handle:         "apostolos",
		Name:           "Apostolos Tzinas",
		Specialization: "smart contract development, algorithms, software engineering",
		Desc: template.HTML(`
			Apostolos is a smart contract software engineer.
			He is pursuing a joint Bachelor and Master's degree in Electrical and Computer Engineering
			at the National Technical University of Athens.
			Apostolos has extensive front-end software engineering experience working at
			<a href="https://www.mayainsights.com/">Maya Insights</a>,
			as well as at NutriDice, where he took on a full-stack engineering role.
			He has worked with numerous programming languages and technical stacks.
		`),
		Image: "apostolos_tzinas.jpg",
	},
	"dionysis": {
		Handle:         "dionysis",
		Name:           "Dr. Dionysis Zindros",
		Specialization: "scaling blockchains, light clients, fast bootstrapping, algorithms, software engineering",
		Desc: template.HTML(`
			Dionysis is a post-doctoral blockchain researcher at Stanford University. He holds a PhD from the University
			of Athens and an Electrical and Computer Engineering degree from the National Technical University of Athens.
			He specializes in blockchain scalability via light clients and fast bootstrapping. Among other venues he has
			published in IEEE S&P (Oakland, ACM CCS, ESORICS, and Financial Crypto, and presented at Black Hat Europe and
			<a href="https://eprint.iacr.org/2017/963.pdf">Non-Interactive Proofs of Proof-of-Work</a>,
			<a href="https://eprint.iacr.org/2018/1239.pdf">Proof-of-Stake Sidechains</a>, and
			<a href="https://eprint.iacr.org/2018/1048.pdf">Proof-of-Work Sidechains</a>.
		`),
		Image: "dionysis_zindros.jpg",
	},
}
