package main

import "html/template"

var Members map[string]TeamMember = map[string]TeamMember{
	"orestis": {
		Handle:         "orestis",
		Name:           "Dr. Orestis Alpos",
		Specialization: "consensus, distributed systems, threshold cryptography, applied cryptography",
		Department: "science",
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
	"dominik": {
		Handle:         "dominik",
		Name:           "Dominik Apel",
		Specialization: "software engineering, auditing, offensive security",
		Department: "engineering",
		Desc: template.HTML(`
			Dominik is a Software Engineer based in Vienna, Austria. He holds a Master's degree in Computer Engineering from the 
			Technical University of Vienna (TU Wien) in Austria. He has multiple years of experience in full-stack software engineering
			as well as some experience in offensive security. Dominik is currently focused on smart contract development
			and source code auditing at Common Prefix.
		`),
		Image: "dominik_apel.png",
	},
	"jason": {
		Handle:         "jason",
		Name:           "Jason Athanasoglou",
		Specialization: "software engineering, interoperability, algorithms, DevOps",
		Department: "engineering",
		Desc: template.HTML(`
			    Jason is a software engineer based in Athens, Greece, with extensive experience in full-stack development, open-source contributions, 
				and a strong focus on building AI-powered applications. 
				He holds an Integrated Master’s degree in Electrical and Computer Engineering from the University of Patras. 
				At Common Prefix, he develops blockchain and Web3 applications, focusing on decentralized and scalable solutions.`),
		Image: "jason_athanasoglou.jpg",
	},
	"lukas": {
		Handle:         "lukas",
		Name:           "Dr. Lukas Aumayr",
		Specialization: "distributed systems, scaling blockchains, blockchain interoperability, applied cryptography, protocol analysis",
		Department: "science",
		Desc: template.HTML(`
            Lukas is a blockchain researcher specializing in distributed systems, applied cryptography,
            and protocol analysis, with a focus on blockchain scalability and interoperability.
            Among others, his work includes advancements in Payment/Virtual Channel Network protocols,
            bridges, light clients, crypto-economic incentives, and BitVM.
            Lukas earned his PhD from the Technical University of Vienna (TU Wien),
            under the supervision of Matteo Maffei and Pedro Moreno-Sanchez.
            He has published in venues such as ACM CCS, USENIX Security, IEEE S&P (Oakland), and NDSS.
		`),
		Image: "lukas_aumayr.jpg",
	},
	"zeta": {
		Handle:         "zeta",
		Name:           "Dr. Zeta Avarikioti",
		Specialization: "distributed systems, scaling blockchains, cryptoeconomic incentive analysis",
		Department: "science",
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

	"joao": {
		Handle:         "joao",
		Name:           "João Pedro Coelho de Azevedo",
		Specialization: "smart contract development, large scale system design, software engineering",
		Department: "engineering",
		Desc: template.HTML(`
			João is a software engineer with experience in blockchain, fintech, and automotive technologies. 
			He holds a Master’s degree in Electronics and Telecommunications Engineering from the University of Aveiro. 
			Over the years he has contributed to multiple startups (notably <a href="https://www.qomodo.me/">Qomodo</a>, <a href="https://www.topos.network/">Topos Network</a> and <a href="https://veniam.com/">Veniam</a>), building financing risk assessment systems, smart contracts, and connected mobility solutions. 
			His work spans backend engineering, smart contract development, and large-scale system architecture.
			João is a certified PMP, IEEE member, and the inventor of 13 granted patents in intelligent networking and automotive communications.
		`),
		Image: "joao_azevedo.jpg",
	},
	"tor": {
		Handle:         "tor",
		Name:           "Tor Bognaes",
		Specialization: "business development, finance, operations",
		Department: "operations",
		Desc: template.HTML(`
			Tor supports the growth of Common Prefix, working across business development, finance, marketing, and operations. 
			His background includes work at Amazon - where he helped scale the EU last-mile logistics network - and at a growth equity investment firm. 
			He holds MSc degrees in Industrial Economics & Technology Management from the Norwegian University of Science and Technology and in Finance from the Norwegian School of Economics.
		`),
		Image: "tor_bognaes.jpg",
	},
	"dejan": {
		Handle:         "dejan",
		Name:           "Dejan Cabrilo",
		Specialization: "software engineering, auditing, quantitative trading",
		Department: "engineering",
		Desc: template.HTML(`
			Dejan is a software engineer based in Belgrade, Serbia. He holds a Bachelor's degree in Computing and Information Systems from Goldsmiths, University of London. 
			He has many years of experience building high-availability full stack software. He has also done work in data science and low-level programming. 
			Previously he worked in an e-commerce startup and a quantitative trading firm in traditional finance.
		`),
		Image: "dejan_cabrilo.png",
	},
	"pyrros": {
		Handle:         "pyrros",
		Name:           "Dr. Pyrros Chaidos",
		Specialization: "zero-knowledge proofs, voting",
		Department: "science",
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
	"james": {
		Handle:         "james",
		Name:           "Dr. James Hsin-yu Chiang",
		Specialization: "applied cryptography, privacy-enhancing technologies, blockchain protocols",
		Department: "science",
		Desc: template.HTML(`
			James is a cryptography postdoctoral researcher at Aarhus University, advised by Ivan Damgård. He holds a PhD from the Technical University of Denmark where he was awarded the DTU Compute Fellowship and advised by Alberto Lluch-Lafuente, Bernardo David and Massimo Bartoletti. 
			He holds a BSc with distinction from UCLA. His research focuses on privacy-preserving techniques motivated by recent blockchain applications and includes practical advancements in <a href="https://eprint.iacr.org/2025/113.pdf">post-quantum threshold ring signatures</a>, novel notions of <a href="https://eprint.iacr.org/2023/943.pdf">differential privacy for secure multi-party computation</a> (MPC), and the first formal model of <a href="https://arxiv.org/pdf/2106.01870">miner-extractable-value</a> (MEV).
			His research has been published in ACM CCS, Financial Cryptography, AFT and recognized with the Sui Academic Research Award.
			Prior to academia, James designed flight hardware for Mars rovers at JPL NASA, advised technology firms on three continents at the Boston Consulting Group, co-founded the leading solar business software company (Eturnity) and contributed to multiple open-source implementations of the original Bitcoin protocol (Core/Libbitcoin).	
		`),
		Image: "james_chiang.jpeg",
	},
	"makis": {
		Handle:         "makis",
		Name:           "Michalis Christou",
		Specialization: "smart contract development, software engineering, bridges",
		Department: "engineering",
		Desc: template.HTML(`
			Michalis Christou is a Software Engineer based in Larnaca, Cyprus. He holds a Master’s degree in Advanced Computing from Imperial College London and a Bachelor’s degree in Computer Engineering from the University of Cyprus. 
			He has multiple years of experience spanning Ethereum smart contract development, embedded systems, and more recently bridges. Michalis is currently focused on smart contract development and bridges at Common Prefix.
		`),
		Image: "michalis_christou.jpg",
	},
	"bernardo": {
		Handle:         "bernardo",
		Name:           "Prof. Bernardo David",
		Specialization: "multiparty computation, consensus",
		Department: "science",
		Desc: template.HTML(`
			Bernardo David is an Associate Professor at the IT University of Copenhagen, working on cryptographic protocols
			for multiparty computation and blockchain consensus/applications with the support of Concordium Foundation,
			Independent Research Fund Denmark and Protocol Labs research grants. He holds a PhD in Computer Science
			from Aarhus University, under the supervision of Ivan Damgård and Jesper Buus Nielsen, and was previously an Assistant Professor at the Tokyo Institute of Technology,
			where his work was supported by a JSPS research grant. Bernardo&rsquo;s research has been published as over 30 articles in
			scientific journals and conferences. Among his main contributions is
			<a href="https://eprint.iacr.org/2016/889.pdf">the first provably secure protocol for proof-of-stake blockchains</a>.
			Besides academic activities, he has been a consultant for a number of industry projects
			on blockchains and information security, having served as scientific advisor to Cardano and Concordium.
		`),
		Image: "bernardo_david.jpg",
	},
	"robin": {
		Handle:         "robin",
		Name:           "Dr. Robin Fritsch",
		Specialization: "DeFi, game theory, graph algorithms",
		Department: "science",
		Desc: template.HTML(`
            Robin holds a PhD from ETH Zurich, advised by Roger Wattenhofer, focusing on game theory and decentralized finance. 
			He completed his MSc degree in Mathematics in Data Science from the Technical University of Munich.  
			Among other venues, he has published in AFT, ACM CCS, and AAMAS. Highlights of his research include the papers <a href="https://arxiv.org/pdf/2201.12303">The Price of Majority Support</a>, 
			<a href="https://arxiv.org/pdf/2206.04634">The Economics of Automated Market Makers</a>, 
			and <a href="https://drops.dagstuhl.de/storage/00lipics/lipics-vol282-aft2023/LIPIcs.AFT.2023.24/LIPIcs.AFT.2023.24.pdf">Batching trades on Automated Market Makers</a>.
		`),
		Image: "robin_fritsch.jpg",
	},
	"lioba": {
		Handle:         "lioba",
		Name:           "Lioba Heimbach",
		Specialization: "DeFi, cryptoeconomic incentives analysis, MEV, game theory",
		Department: "science",
		Desc: template.HTML(`
			Lioba is a doctoral student at ETH Zurich advised by Roger Wattenhofer. 
			She holds a MSc degree in Electrical Engineering and Information Technology from ETH Zurich. 
			Her research focuses on empirical measurements and the analysis of cryptoeconomic incentives in the blockchain ecosystem with a focus on decentralized finance. 
			Among other venues, she has published in IEEE S&P (Oakland), FC, AFT, AsiaCCS, IMC, and DISC. 
			Highlights of her research include the papers <a href="https://arxiv.org/abs/2401.01622">Non-Atomic Arbitrage in Decentralized Finance</a>, <a href="https://arxiv.org/abs/2203.11520">SoK: Preventing Transaction Reordering Manipulations in Decentralized Finance</a>, and  <a href="https://arxiv.org/pdf/2304.11478">Base Fee Manipulation In Ethereum's EIP-1559 Transaction Fee Mechanism</a>.
		`),
		Image: "lioba_heimbach.jpg",
	},
	"nikolas": {
		Handle:         "nikolas",
		Name:           "Nikolaos Kamarinakis",
		Specialization: "smart contract development, auditing, software engineering, offensive security",
		Department: "engineering",
		Desc: template.HTML(`
			Nikolas is a software engineer based in Athens, Greece.
			He holds a Bachelor&rsquo;s degree in Computer Science with a minor in Cybersecurity from the University of Maryland.
			He has multiple years of experience in full-stack software engineering and open-source development,
			as well as some experience in offensive security.
			Nikolas is currently focused on smart contract development and auditing at Common Prefix.
		`),
		Image: "nikolaos_kamarinakis.jpg",
	},
	"dimitris": {
		Handle:         "dimitris",
		Name:           "Dr. Dimitris Karakostas",
		Specialization: "cryptocurrency wallets, macroeconomics, checkpointing",
		Department: "science",
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
	"haris": {
		Handle:         "haris",
		Name:           "Haris Karavasilis",
		Specialization: "business development, project management, quantitative finance, risk management",
		Department: "operations",
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
	"odysseas": {
		Handle:         "odysseas",
		Name:           "Odysseas Sofikitis",
		Specialization: "consensus, proof-of-work, machine learning",
		Department: "science",
		Desc: template.HTML(`
		Odysseas is a blockchain researcher based in Athens,
		specializing in consensus mechanisms and proof-of-work protocols for blockchain systems.
		He holds a Master's degree in Electrical and Computer Engineering from Aristotle University of Thessaloniki,
		with a focus on inverse rendering and machine learning.
		He previously worked at the Information Technologies Institute,
		where he conducted research on GANs and image-to-image translation.
		`),
		Image: "odysseas_sofikitis.jpg",
	},
	"julian": {
		Handle:         "julian",
		Name:           "Prof. Julian Loss",
		Specialization: "cryptography, distributed systems, consensus",
		Department: "science",
		Desc: template.HTML(`
		Julian is a professor at CISPA Helmholtz Center for Information Security, focusing on cryptography and its applications to distributed algorithms.  
		He holds a PhD in Mathematics from Ruhr University Bochum, advised by Eike Kiltz, and a MSc in Computer Science from ETH Zurich.  
		Julian was a postdoctoral researcher at the University of Maryland and Carnegie Mellon University. 
		His interests include provable security, particularly digital signature schemes and algorithms for distributed consensus.
		Among other venues, he has published in ACM CCS, ASIACRYPT, EUROCRYPT, FC, and IEEE S&P (Oakland). 
		Highlights of his research include the papers <a href="https://eprint.iacr.org/2017/620.pdf"> The Algebraic Group Model and its Applications</a>, 
		<a href="https://infoscience.epfl.ch/server/api/core/bitstreams/3820a5e9-3807-48ae-9604-c37037cc907d/content">On the Security of Two-Round Multi-Signatures</a>, 
		and <a href="https://eprint.iacr.org/2020/945.pdf">On the (in)security of ROS</a>.
		`),
		Image: "julian_loss.jpg",
	},
	"orfeas": {
		Handle:         "orfeas",
		Name:           "Dr. Orfeas Stefanos Thyfronitis Litos",
		Specialization: "distributed systems, scaling blockchains, interoperability, protocol analysis, cryptoeconomic incentive analysis, voting, software engineering, payment channels",
		Department: "science",
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
	"jakov": {
		Handle:         "jakov",
		Name:           "Jakov Mitrovski",
		Specialization: "smart contract development, auditing, software engineering, algorithms, complexity",
		Department: "engineering",
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
		Department: "engineering",
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
		Department: "science",
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
	"nikos": {
		Handle:         "nikos",
		Name:           "Nikos Sfakianakis",
		Specialization: "software engineering, consensus, offensive security",
		Department: "engineering",
		Desc: template.HTML(`Nikos is a software engineer based in Athens. He is currently wrapping up his Master’s Degree in Electrical and Computer Engineering at the National Technical University of Athens where he is writing his thesis on Consensus Algorithms. 
		He has previously worked at Mysten Labs as a Blockchain Engineer and completed an internship at the Cybersecurity Department of OTE Group. He has also competed in several Ethical Hacking and Offensive Security competitions.
		`),
		Image: "nikos_sfakianakis.jpg",
	},	
	"semeli": {
		Handle:         "semeli",
		Name:           "Semeli Spanou",
		Specialization: "operations, logistics, biotech, data analytics",
		Department: "operations",
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
		Department: "engineering",
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
		Image: "apostolos_tzinas.png",
	},
	"ristic": {
		Handle:         "ristic",
		Name:           "Nikola Ristić",
		Specialization: "smart contract development, full-stack web development, software engineering, algorithms, circus acrobatics",
		Department: "engineering",
		Desc: template.HTML(`
			Nikola is a self-taught software engineer based in Belgrade, Serbia. He started with algorithms in elementary school and won 2 gold medals in national high-school competitions in Serbia. Before his web3 journey he has helped various startups scale their products and teams, notably <a href="https://hellosuper.com">Super</a> and <a href="https://www.amondo.com">Amondo</a>, and has 10 years of freelance experience as a full-stack web developer.

		`),
		Image: "nikola_ristic.jpg",
	},
	"nikos_vlastaras": {
		Handle:         "nikos_vlastaras",
		Name:           "Nikos Vlastaras",
		Specialization: "software engineering, distributed systems",
		Department: "engineering",
		Desc: template.HTML(`
			Nikos is a software engineer based in Athens with a degree from the National Technical University of Athens. 
			Before joining Common Prefix he worked at Panther Labs, where he contributed to production software and large-scale engineering projects. 
			At Common Prefix he applies his experience in distributed systems to help build robust and scalable infrastructure.
		`),
		Image: "nikos_vlastaras.jpg",
	},
	"dionysis": {
		Handle:         "dionysis",
		Name:           "Dr. Dionysis Zindros",
		Specialization: "consensus, light clients, bridges, interoperability, fast bootstrapping, algorithms, software engineering",
		Department: "science",
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
