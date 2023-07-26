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
			Previously, he worked at <a href="https://www.paraswap.io/">ParaSwap</a> where he was responsible for developing a significant part of the core aggregation algorithm.
			He is interested in Cryptography, Security, Consensus Protocols, Decentralised Finance, and Ethereum.
			One of Shresth&rsquo;s research highlights is the paper
			<a href="https://arxiv.org/pdf/2209.08673.pdf">Proofs of Proof-of-Stake with Sublinear Complexity</a>.
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
			Among other venues, she has published in Financial Crypto, AFT and SODA. Highlights of her research include the papers
			<a href="https://arxiv.org/pdf/1905.11360.pdf">BRICK: Asynchronous Payment Channels</a>,
			<a href="https://eprint.iacr.org/2019/1092.pdf">Cerberus Channels: Incentivizing Watchtowers for Bitcoin</a>, and
			<a href="https://arxiv.org/pdf/1910.10434.pdf">Divide and Scale: Formalization of Distributed Ledger Sharding Protocols</a>.
		`),
		Image: "zeta_avarikioti.jpg",
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
			voting in particular. Among other venues, he has published in EUROCRYPT, ACM CCS, PKC, ESORICS, and the Journal
            of Cryptology. Highlights of his research include the papers
            <a href="https://eprint.iacr.org/2021/916.pdf">Mithril: Stake-based Threshold Multisignatures</a>,
            <a href="https://eprint.iacr.org/2016/263.pdf">Efficient Zero-Knowledge Arguments for Arithmetic Circuits in
            the Discrete Log Setting</a>, and
            <a href="https://eprint.iacr.org/2016/368.pdf">Foundations of Fully Dynamic Group Signatures</a>.
		`),
		Image: "pyrros_chaidos.jpg",
	},
	"bernardo": {
		Handle:         "bernardo",
		Name:           "Prof. Bernardo David",
		Specialization: "multiparty computation, consensus",
		Desc: template.HTML(`
            Bernardo David is an Associate Professor at the IT University of Copenhagen, working on cryptographic protocols
            for multiparty computation and blockchain consensus/applications with the support of Concordium Foundation,
            Independent Research Fund Denmark and Protocol Labs research grants. He holds a Ph.D. in Computer Science
            from Aarhus University and was previously an Assistant Professor at the Tokyo Institute of Technology,
            where his work was supported by a JSPS research grant. Bernardo&rsquo;s research has been published as over 30 articles in
            scientific journals and conferences. Among his main contributions is
			<a href="https://eprint.iacr.org/2016/889.pdf">the first provably secure protocol for proof-of-stake blockchains</a>.
			Besides academic activities, he has been a consultant for a number of industry projects
            on blockchains and information security, having served as scientific advisor to Cardano and Concordium.
		`),
		Image: "bernardo_david.jpg",
	},
	"nikolas": {
		Handle:         "nikolas",
		Name:           "Nikolaos Kamarinakis",
		Specialization: "smart contract development, testing, software engineering, offensive security",
		Desc: template.HTML(`
			Nikolas is a software engineer based in Athens, Greece and Maryland, USA.
			He is currently pursuing a bachelor&rsquo;s degree in Computer Science at the University of Maryland,
			specializing in Cybersecurity.
			He has multiple years of experience in full-stack software engineering and open-source development,
			as well as some expierience in offensive security.
			Nikolas is currently focused on smart contract development at Common Prefix
			as well as at the <a href="https://researchdao.io">Research DAO</a>,
			where he also coordinates operations.
		`),
		Image: "nikolaos_kamarinakis.jpg",
	},
	"haris": {
		Handle:         "haris",
		Name:           "Haris Karavasilis",
		Specialization: "business development, project management, quantitative finance, risk management",
		Desc: template.HTML(`
            Haris works on the business side of things at Common Prefix.
            His background includes roles at Amazon and Piraeus Bank,
            where he gained valuable experience in operations, finance and risk management.
            He holds a Master's degree in Quantitative Finance and Risk Management from Bocconi University
            and an Electrical & Electronic Engineering Bachelor's degree from the University of Manchester.
            At Common Prefix, he contributes to both business development and project management,
            aiming to improve operational and financial efficiency.
		`),
		Image: "haris_karavasilis.png",
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
			presented at Black Hat Europe and Asia. Highlights of his research include the papers
            <a href="https://eprint.iacr.org/2019/034.pdf">A Formal Treatment of Hardware Wallets</a>,
            <a href="https://arxiv.org/pdf/1907.02434.pdf">Cryptocurrency Egalitarianism: A Quantitative Approach</a>, and
            <a href="https://eprint.iacr.org/2020/173.pdf">Securing Proof-of-Work Ledgers via Checkpointing</a>.
		`),
		Image: "dimitris_karakostas.jpg",
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
		Specialization: "smart contract development, large scale system design, software engineering, consensus, DevOps",
		Desc: template.HTML(`
			Dimitris is a software engineer based in Thessaloniki. He holds a Bachelor&rsquo;s degree in Computer Science from Aristotle
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
			He is pursuing a Master&rsquo;s degree in Electrical and Computer Engineering
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
			He is pursuing a joint Bachelor and Master&rsquo;s degree in Electrical and Computer Engineering
			at the National Technical University of Athens.
			Apostolos has extensive front-end software engineering experience working at
			<a href="https://www.mayainsights.com/">Maya Insights</a>,
			as well as at NutriDice, where he took on a full-stack engineering role.
			He has worked with numerous programming languages and technical stacks.
			One of Apostolos&rsquo;s research highlights is the paper
			<a href="https://eprint.iacr.org/2023/605.pdf">The Principal–Agent Problem in Liquid Staking</a>,
			published in Financial Cryptography 2023&rsquo;s 7th Workshop on Trusted Smart Contracts (WTSC).
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
			published in IEEE S&P (Oakland, ACM CCS, ESORICS, and Financial Crypto, and presented at Black Hat Europe and Asia. Highlights of his research include the papers
			<a href="https://eprint.iacr.org/2017/963.pdf">Non-Interactive Proofs of Proof-of-Work</a>,
			<a href="https://eprint.iacr.org/2018/1239.pdf">Proof-of-Stake Sidechains</a>, and
			<a href="https://eprint.iacr.org/2018/1048.pdf">Proof-of-Work Sidechains</a>.
		`),
		Image: "dionysis_zindros.jpg",
	},
}
