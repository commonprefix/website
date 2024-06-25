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
	"orestis": {
		Handle:         "orestis",
		Name:           "Dr. Orestis Alpos",
		Specialization: "consensus, distributed systems, threshold cryptography, applied cryptography",
		Desc: template.HTML(`
			Orestis is a blockchain researcher based in Bern, Switzerland.
			He holds a PhD from the University of Bern, advised by Christian Cachin,
			and a Computer Engineering MSc degree from the National Technical University of Athens.
			He specializes in distributed systems and applied cryptography.
			His publications include protocols for
			<a href="https://eprint.iacr.org/2023/1103.pdf">asynchronous PoS consensus</a>,
			<a href="https://arxiv.org/pdf/2307.02954.pdf">front-running prevention</a>,
			<a href="https://arxiv.org/pdf/2006.04616.pdf">consensus</a> and
			<a href="https://eprint.iacr.org/2022/1767.pdf">cryptography</a>
			with flexible trust assumptions, as well as results on
			<a href="https://arxiv.org/pdf/2107.11331.pdf">composing distributed systems</a>.
		`),
		Image: "orestis_alpos.jpg",
	},
	"zeta": {
		Handle:         "zeta",
		Name:           "Dr. Zeta Avarikioti",
		Specialization: "distributed systems, scaling blockchains, cryptoeconomic incentive analysis",
		Desc: template.HTML(`
			Zeta is an upcoming Assistant Professor at the Technical University of Vienna (TU Wien) in Austria.
			She is currently a post-doctoral blockchain researcher at TU Wien working with professor Matteo Maffei.
			She graduated with a PhD from ETH Zürich, advised by Roger Wattenhofer, and holds an engineering degree from the
			National Technical University of Athens and a masters degree from National Kapodistian University of Athens.
			She specializes in distributed systems, scaling blockchains via sharding and channels,
			and the analysis of cryptoeconomic incentives.
			Among other venues, she has published in USENIX Security, CSF, Financial Cryptography, AFT, AAAI, and SODA.
			Highlights of her research include the papers
			<a href="https://arxiv.org/pdf/1905.11360.pdf">BRICK: Asynchronous Payment Channels</a>,
			<a href="https://eprint.iacr.org/2019/1092.pdf">Cerberus Channels: Incentivizing Watchtowers for Bitcoin</a>, and
			<a href="https://arxiv.org/pdf/1910.10434.pdf">Divide and Scale: Formalization of Distributed Ledger Sharding Protocols</a>.
		`),
		Image: "zeta_avarikioti.jpg",
	},
	"alexis": {
		Handle:         "alexis",
		Name:           "Alexis Brezas",
		Specialization: "smart contract development, algorithms, machine learning, IoT, DevOps, software engineering",
		Desc: template.HTML(`
			Alex is a proficient software engineer with over 15 years of experience in the tech industry.
			In his web2 life, he has served as CTO at
			<a href="https://www.mayainsights.com/">Maya Insights</a>,
			a data warehousing solution, and
			<a href="https://geekbot.com/">Geekbot</a>,
			a popular Slack chatbot, as well as various engineering roles in diverse projects such as
			<a href="https://www.e-food.gr/">a popular Greek food delivery service</a>,
			<a href="https://en.wikipedia.org/wiki/Zino">a social network</a>,
			a 3D game, ecommerce solutions, and
            <a href="https://balena.io">IoT</a>.
			At Common Prefix, he focuses on blockchain engineering and smart contract development.
		`),
		Image: "alexis_brezas.jpg",
	},
	"pyrros": {
		Handle:         "pyrros",
		Name:           "Dr. Pyrros Chaidos",
		Specialization: "zero-knowledge proofs, voting",
		Desc: template.HTML(`
			Pyrros is a post-doctoral blockchain researcher at the University of Athens. He holds a PhD from the
			University College of London, advised by Jens Groth and Allan Sikk, an MSc in Information Security from the University College of London, an MSc in
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
			from Aarhus University, under the supervision of Ivan Damgård and Jesper Buus Nielsen, and was previously an Assistant Professor at the Tokyo Institute of Technology,
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
			He holds a Bachelor&rsquo;s degree in Computer Science with a minor in Cybersecurity from the University of Maryland.
			He has multiple years of experience in full-stack software engineering and open-source development,
			as well as some experience in offensive security.
			Nikolas is currently focused on smart contract development at Common Prefix.
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
			Aggelos Kiayias. He holds a PhD from the University of Edinburgh, advised by Aggelos Kiayias, and an Electrical and Computer Engineering
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
		Orfeas is currently a research associate at Imperial College London.
		He obtained his
		<a
			href="https://gitlab.com/orfeasLitos/thesis/-/raw/master/thesis.pdf?inline=false"
		>PhD in Cryptography and Blockchains</a> at the University of Edinburgh in 2021, under the supervision of Aggelos Kiayias. He has worked on <a href="http://fc17.ifca.ai/preproceedings/paper_37.pdf">building and analyzing</a> decentralized applications on blockchains,
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
			Dimitris is a software engineer based in Athens, Greece.
			He holds a Bachelor's degree in Computer Science
			from Aristotle University of Thessaloniki.
			At Common Prefix, he is a software engineer focusing on
			Solidity smart contracts and basic consensus development.
			Besides his web3 expertise, Dimitris has significant experience
			in building and scaling web applications.
			Before joining Common Prefix, he worked at
			<a href="https://geekbot.com/">Geekbot</a>,
			one of the most popular Slack bots, and
			<a href="https://www.amondo.com/">Amondo</a>,
			a social media startup with over a million users,
			where he led the infrastructure team.
		`),
		Image: "dimitris_lamprinos.jpg",
	},
	"jakov": {
		Handle:         "jakov",
		Name:           "Jakov Mitrovski",
		Specialization: "smart contract development, auditing, software engineering, algorithms, complexity",
		Desc: template.HTML(`
			Jakov is a software engineer based in Munich,
			currently pursuing his Master's degree in Informatics at the Technical University of Munich.
			He has earned his Bachelor's degree in Computer Science
			from the University Cyril and Methodius in Skopje.
			Jakov specializes in blockchain engineering, algorithms and complexity, and auditing.
			He has worked as a cloud engineer at
			<a href="https://www.etas.com/">ETAS (Bosch) GmbH</a>
			and has experience working as a backend and blockchain engineer through various internships at
			<a href="https://www.netcetera.com/">Netcetera</a>,
			<a href="https://www.sorsix.com/">Sorsix</a>, and
			<a href="https://www.loka.com/">LOKA</a>,
			as well as through freelance projects.
		`),
		Image: "jakov_mitrovski.jpg",
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
	"giulia": {
		Handle:         "giulia",
		Name:           "Giulia Scaffino",
		Specialization: "scaling blockchains, blockchain interoperability, cryptoeconomic incentive analysis",
		Desc: template.HTML(`
			Giulia is a doctoral blockchain researcher at the Technical University of Vienna (TU Wien),
			working in the Security and Privacy group with Prof. Matteo Maffei.
			She holds a Master's degree in Nuclear Physics from the University of Pavia, Italy,
			and she has worked as a Salesforce developer at Deloitte Digital in Milan.
			She specializes in blockchain interoperability and scalability protocols,
			light clients, layer-2 solutions, and crypto-economic incentives.
			Among other venues, she has published in the USENIX Security Symposium.
		`),
		Image: "giulia_scaffino.png",
	},
	"semeli": {
		Handle:         "semeli",
		Name:           "Semeli Spanou",
		Specialization: "operations, logistics, biotech, data analytics",
		Desc: template.HTML(`
			Semeli is a member of the Operations team at Common Prefix,
			where she collaborates with a diverse range of stakeholders
			to drive operational efficiency and effectiveness.
			She holds a Master's degree in Biotechnology,
			from the Agricultural University of Athens,
			with a focus on bioinformatics and data analysis.
			Semeli's organisational and analytical skills contribute to
			enhancing operational procedures, streamline reporting,
			and strengthening our logistics management.
		`),
		Image: "semeli_spanou.jpg",
	},
	"apostolos": {
		Handle:         "apostolos",
		Name:           "Apostolos Tzinas",
		Specialization: "smart contract development, algorithms, software engineering, consensus",
		Desc: template.HTML(`
			Apostolos is a blockchain researcher and engineer at Common Prefix,
			specialising in blockchain consensus and decentralised finance.
			His research highlights include
			<a href="https://eprint.iacr.org/2023/1648.pdf">On-Chain Timestamps Are Accurate</a>,
			published in Financial Cryptography 2024, and
			<a href="https://eprint.iacr.org/2023/605.pdf">The Principal–Agent Problem in Liquid Staking</a>,
			published in Financial Cryptography 2023&rsquo;s 7th Workshop on Trusted Smart Contracts (WTSC).
			Apostolos also has extensive experience in deploying and managing both validators and full-nodes
			across the Ethereum and Cosmos ecosystems.
			In the past, as a web engineer at
			<a href="https://www.mayainsights.com/">Maya Insights</a> and NutriDice,
			he has gained extensive experience with a wide range of programming languages and technical stacks.
			Apostolos has a background in algorithms, having competed at national and balkan olympiads in informatics.
		`),
		Image: "apostolos_tzinas.jpg",
	},
	"ristic": {
		Handle:         "ristic",
		Name:           "Nikola Ristić",
		Specialization: "smart contract development, full-stack web development, software engineering, algorithms, circus acrobatics",
		Desc: template.HTML(`
			Nikola is a self-taught software engineer based in Belgrade, Serbia. He started with algorithms in elementary school and won 2 gold medals in national high-school competitions in Serbia. Before his web3 journey he has helped various startups scale their products and teams, notably <a href="https://hellosuper.com">Super</a> and <a href="https://www.amondo.com">Amondo</a>, and has 10 years of freelance experience as a full-stack web developer.

		`),
		Image: "nikola_ristic.jpg",
	},
	"dionysis": {
		Handle:         "dionysis",
		Name:           "Dr. Dionysis Zindros",
		Specialization: "consensus, light clients, bridges, interoperability, fast bootstrapping, algorithms, software engineering",
		Desc: template.HTML(`
			Dionysis is a co-founder and researcher at Common Prefix
			focusing on consensus, light clients, bridges, interoperability, and fast bootstrapping.
			He did his post-doc at Stanford University, advised by David Tse.
			He holds a PhD from the University of Athens, advised by Aggelos Kiayias,
			and an Electrical and Computer Engineering degree from the National Technical University of Athens.
			Among other venues, he has published in IEEE S&P (Oakland), ACM CCS, ESORICS, and Financial Crypto, and
			presented at Black Hat Europe and Asia.
			Highlights of his research include the papers
			<a href="https://eprint.iacr.org/2017/963.pdf">Non-Interactive Proofs of Proof-of-Work</a>,
			<a href="https://eprint.iacr.org/2018/1239.pdf">Proof-of-Stake Sidechains</a>, and
			<a href="https://eprint.iacr.org/2018/1048.pdf">Proof-of-Work Sidechains</a>.
		`),
		Image: "dionysis_zindros.jpg",
	},
}
