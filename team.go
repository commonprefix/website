package main

import "html/template"

// Titles for Leadership
var Titles = map[string]string{
	"dionysis": "CEO & Co-founder",
	"nikolas": "CTO & Co-founder",
	"zeta": "Chief Scientist & Co-founder",
	"haris":   "COO",
	"themis":   "Engineer Lead",
	"apostolos":   "Product Lead & Co-founder",
	"dimitris":   "Co-founder",
}

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
			Lukas is a blockchain researcher 
			<a href="https://dblp.org/pid/265/9073.html">specializing</a> in distributed systems, applied cryptography, and protocol analysis, 
			with a focus on blockchain scalability and interoperability. 
			Among others, his work includes advancements in Payment/Virtual Channel Network protocols, bridges, light clients, 
			crypto-economic incentives, and BitVM. 
			Notably, he has co-authored several papers on light clients and bridges, including 
			<a href="https://eprint.iacr.org/2024/692.pdf">Blink: An Optimal Proof of Proof-of-Work</a>,
			<a href="https://eprint.iacr.org/2024/197.pdf">Alba: The Dawn of Scalable Bridges for Blockchains</a>,
			<a href="https://eprint.iacr.org/2022/1721.pdf">Glimpse: On-Demand PoW Light Client with Constant-Size Storage for DeFi</a>,
			<a href="https://eprint.iacr.org/2025/1158.pdf">BitVM2: Bridging Bitcoin to Second Layers</a>.
			Lukas is a research associate at the University of Edinburgh hosted by Aggelos Kiayias. 
			Lukas holds a PhD from the Technical University of Vienna (TU Wien), where he was advised by Matteo Maffei 
			and Pedro Moreno-Sanchez. His work has been published in top-tier venues such as 
			ACM CCS, USENIX Security, IEEE Security & Privacy, NDSS, FC, and Asiacrypt.		
		`),
		Image: "lukas_aumayr.jpg",
	},
	"zeta": {
		Handle:         "zeta",
		Name:           "Prof. Zeta Avarikioti",
		Specialization: "distributed systems, scaling blockchains, cryptoeconomic incentive analysis",
		Department: "leadership",
		Desc: template.HTML(`
			Zeta is an Assistant Professor at the Technical University of Vienna (TU Wien), where she leads the
			<a href="https://cysec.wien/initiatives/blockchain/">Blockchain Hub</a>, within the university’s Cybersecurity Center. 
			She holds a PhD from ETH Zürich, advised by Roger Wattenhofer, and has previously held research positions at IST Austria and Columbia University.
			She <a href="https://dblp.org/pid/183/6375.html">specializes</a>, in distributed systems, blockchain scalability and interoperability, and the analysis of cryptoeconomic incentives 
			with numerous publications, spanning payment channels (e.g.,
			<a href="https://dl.acm.org/doi/abs/10.1007/978-3-662-64331-0_11">Brick</a>,
			<a href="https://link.springer.com/chapter/10.1007/978-3-030-51280-4_19">Cerberus</a>,
			<a href="https://dl.acm.org/doi/10.1145/3658644.3670373">Crab</a>,
			light clients and bridges (e.g.,		
			<a href="https://eprint.iacr.org/2024/692.pdf">Blink</a>,
			<a href="https://eprint.iacr.org/2024/197.pdf">Alba</a>,
			<a href="https://eprint.iacr.org/2022/1721.pdf">Glimpse</a>,
			<a href="https://eprint.iacr.org/2025/1158.pdf">BitVM</a>),
			consensus (e.g., <a href="https://arxiv.org/pdf/2009.02235">FnF-BFT</a>, <a href="https://arxiv.org/abs/2503.16783">CoBRA</a>), 
			and foundational research on
			<a href="https://link.springer.com/chapter/10.1007/978-3-031-32733-9_10">sharding</a>,
			<a href="https://arxiv.org/abs/2504.14965">layer 2 protocols</a>, and
			<a href="https://arxiv.org/abs/2504.18214">compositional game theory</a>.
			Her work has appeared at leading venues including USENIX Security, ACM CCS, NDSS, CSF, FC, AFT, AAAI, and SODA, 
			and is supported by several major national funding bodies (FWF, WWTF) and industry-funded research donations (Ark Labs, Sui Foundation). 
			She currently serves on the program committees of top-tier conferences such as CCS, USENIX Security, FC, 
			and as Program Chair of the Advances in Financial Technologies (AFT 2025) conference. 
			In 2025, she received the <a href="https://www.ots.at/presseaussendung/OTS_20251016_OTS0003/tu-wien-forscherin-georgia-avarikioti-erhaelt-hedy-lamarr-preis-2025">Hedy Lamarr Prize</a> from the City of Vienna for her contributions to information technology, 
			as well as the <a href="https://brd.chaincode.com/#prize">Bitcoin Research Prize</a>, along with her collaborators, for their work on BitVM.
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
			Dejan is a senior software engineer with decades of C++ experience. He has worked on high-speed, high-concurrency, and high-availability trading systems at hedge funds. 
			He has previously served as CTO and Director of Engineering in multiple companies. 
			He holds a Computer Science degree from the Goldsmiths University of London.
		`),
		Image: "dejan_cabrilo.png",
	},
	"pyrros": {
		Handle:         "pyrros",
		Name:           "Prof. Pyrros Chaidos",
		Specialization: "zero-knowledge proofs, voting",
		Department: "science",
		Desc: template.HTML(`
			Pyrros holds a PhD from UCL, advised by Jens Groth, the person behind modern ZK proofs. 
			He <a href="https://dblp.org/pid/117/9032.html">specializes</a> in zero-knowledge proof systems, with applications on proof-of-stake blockchains 
			and voting in particular, with notable scientific contributions including
			<a href="http://www0.cs.ucl.ac.uk/staff/P.Chaidos/fosad.pdf">Efficient Zero-Knowledge Proof Systems</a>,
			<a href="https://eprint.iacr.org/2016/263.pdf">Efficient Zero-Knowledge Arguments for Arithmetic Circuits in the Discrete Log Setting</a>,
			<a href="https://eprint.iacr.org/2021/916.pdf">Mithril: Stake-based Threshold Multisignatures</a>.			
			Additionally, he has audited multiple deployed zero-knowledge and cryptographic codebases.
		`),
		Image: "pyrros_chaidos.jpg",
	},
	"james": {
		Handle:         "james",
		Name:           "Dr. James Hsin-yu Chiang",
		Specialization: "applied cryptography, privacy-enhancing technologies, blockchain protocols",
		Department: "science",
		Desc: template.HTML(`
			James is a postdoctoral researcher at the Systems Secruity Group at ETH Zurich, advised by Srdjan Capkun.
			Previously, he was a cryptography postdoctoral researcher at Aarhus University, advised by Ivan Damgård.
			He holds a PhD from the Technical University of Denmark where he was awarded the DTU Compute Fellowship and advised by Alberto Lluch-Lafuente, Bernardo David and Massimo Bartoletti. 
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
			Bernardo is an Associate Professor <a href="https://dblp.org/pid/15/10106.html">working on</a> cryptographic protocols for multiparty computation (MPC) and blockchains at the IT-University of Copenhagen.
			Among his main contributions is <a href="https://eprint.iacr.org/2016/889.pdf">the first provably secure protocol for proof-of-stake blockchains (Ouroboros)</a>
			and <a href="https://eprint.iacr.org/2023/1651.pdf">the first publicly verifiable secret sharing (PVSS) with linear complexity (SCRAPE).</a>
			He has also worked on privacy preserving (auditable) smart contracts and cryptocurrencies, random beacons, sharding, signatures, theshold cryptography, 
			time-based cryptography and zero knowledge. 
			His work has been published in top-tier venues such as Asiacrypt, Crypto, Eurocrypt, CCS and FC. 
			Before working with Common Prefix, he has also consulted for IOG (Cardano) and Concordium.
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
	"nikolas": {
		Handle:         "nikolas",
		Name:           "Nikolaos Kamarinakis",
		Specialization: "interoperability, smart contracts, auditing, software engineering, offensive security",
		Department: "leadership",
		Desc: template.HTML(`
		Nikolas is the CTO of Common Prefix, specializing in blockchain interoperability. 
		He has represented Greece at the European Cyber Security Challenge, three years in a row. 
		He holds a Bachelor's degree in Computer Science from the University of Maryland with a specialization in Cybersecurity. 
		He has multiple years of experience in full-stack software engineering, open-source development, and offensive security. 
		Among various projects, Nikolas led the <a href="https://www.commonprefix.com/clients/axelar">Axelar<>XRPL integration</a>
		and the end-to-end development of an EVM-to-EVM multisig-based bridge, which have processed over $200M in volume.
		`),
		Image: "nikolaos_kamarinakis.jpg",
	},
	"dimitris": {
		Handle:         "dimitris",
		Name:           "Dr. Dimitris Karakostas",
		Specialization: "cryptocurrency wallets, macroeconomics, checkpointing",
		Department: "science",
		Desc: template.HTML(`
			Dimitris is a cryptography researcher specialized in blockchain-based distributed ledger systems. 
			His primary focus is on protocol design and analysis, game theory, and macroeconomic aspects of cryptocurrency systems. 
			Notably, he was a core contributor of the Edinburgh Decentralization Index (EDI), where he worked on defining and measuring the notion of decentralization in blockchain systems. 
			Dimitris holds a PhD from the University of Edinburgh, where he was advised by Aggelos Kiayias. 
			His work has been published in venues like ACM CCS, Financial Cryptography, ESORICS, and ACNS.
		`),
		Image: "dimitris_karakostas.jpg",
	},
	"haris": {
		Handle:         "haris",
		Name:           "Haris Karavasilis",
		Specialization: "business development, project management, quantitative finance, risk management",
		Department: "leadership",
		Desc: template.HTML(`
			Haris is the COO of Common Prefix. His background includes roles at
			Amazon and Piraeus Bank, where he gained valuable experience in operations, finance, and
			risk management. He holds a Master’s degree in Quantitative Finance and Risk Management
			from Bocconi University and an Electrical & Electronic Engineering Bachelor’s degree from
			the University of Manchester. At Common Prefix, he leads multiple facets of our business,
			including finance, legal, sales & marketing.
		`),
		Image: "haris_karavasilis.jpg",
	},
	"myrto": {
		Handle:         "myrto",
		Name:           "Myrsini Koulouri",
		Specialization: "legal, operations",
		Department: "operations",
		Desc: template.HTML(`
			Myrsini is a member of the Operations team at Common Prefix, where she supports the legal and regulatory aspects of the company's operational strategy. 
			She holds a Law degree from Aristotle University of Thessaloniki and is a member of the Athens Bar Association. 
			Her legal expertise and analytical mindset help ensure regulatory compliance, streamline internal processes, and reinforce the legal foundations of the organization’s activities.
		`),
		Image: "myrsini_koulouri.jpeg",
	},
	"odysseas": {
		Handle:         "odysseas",
		Name:           "Odysseas Sofikitis",
		Specialization: "consensus, proof-of-work, machine learning",
		Department: "science",
		Desc: template.HTML(`
		Odysseas is a blockchain researcher specializing in consensus protocols. 
		He holds an MEng in Electrical and Computer Engineering from Aristotle University of Thessaloniki and has previously worked on
		designing consensus mechanisms, “consensus-less” architectures, Byzantine fault-tolerant systems, accountable protocols and 
		blockchain interoperability protocols such as light clients. 
		He is co-author of the paper <a href="https://arxiv.org/pdf/2501.14931">Pod: An Optimal-Latency, Censorship-Free, and Accountable Generalized Consensus Layer</a>.
		`),
		Image: "odysseas_sofikitis.jpg",
	},
	"julian": {
		Handle:         "julian",
		Name:           "Prof. Julian Loss",
		Specialization: "cryptography, distributed systems, consensus",
		Department: "science",
		Desc: template.HTML(`
		Julian is a tenured Professor at the CISPA Helmholtz Center for Information Security. 
		He earned his Ph.D. in Mathematics from Ruhr University Bochum under the supervision of Eike Kiltz and subsequently held 
		postdoctoral positions at the University of Maryland with Jonathan Katz and at Carnegie Mellon University with Elaine Shi. 
		His <a href="https://dblp.org/pid/184/3870.html">research</a> focuses on the intersection of cryptography and distributed protocols, 
		with particular expertise in digital signatures, payments, consensus, and randomness generation.
		His work has been published in top-tier conferences, including IEEE S\&P, ACM CCS, CRYPTO, PODC, FC, AsiaCCS, and TCC.
		`),
		Image: "julian_loss.jpg",
	},
	"orfeas": {
		Handle:         "orfeas",
		Name:           "Dr. Orfeas Stefanos Thyfronitis Litos",
		Specialization: "distributed systems, scaling blockchains, interoperability, protocol analysis, cryptoeconomic incentive analysis, voting, software engineering, payment channels",
		Department: "science",
		Desc: template.HTML(`
		Orfeas is a research associate at Imperial College London. 
		He obtained his <a href="https://gitlab.com/orfeasLitos/thesis/-/raw/master/thesis.pdf?inline=false">PhD in Cryptography and Blockchains</a> 
		at the University of Edinburgh in 2021, under the supervision of Aggelos Kiayias. 
		He has <a href="https://dblp.org/pid/197/1429.html">worked</a> on <a href="http://fc17.ifca.ai/preproceedings/paper_37.pdf">building and analyzing</a> decentralized applications on blockchains,
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
			Jakov is a software engineer and auditor based in Skopje. 
			He holds a Master's degree in Informatics from the Technical University of Munich. 
			Jakov specializes in the design of blockchain systems, protocol development, algorithms, and complexity, and in auditing. 
			He has multi-year experience in software engineering, gained by working as a cloud engineer at <a href="https://www.etas.com/">ETAS (Bosch) GmbH</a> 
			and as a backend and blockchain engineer at <a href="https://www.netcetera.com/">Netcetera</a>, <a href="https://www.sorsix.com/">Sorsix</a>, and <a href="https://www.loka.com/">LOKA</a>.
			`),
		Image: "jakov_mitrovski.jpg",
	},
	"themis": {
		Handle:         "themis",
		Name:           "Themis Papameletiou",
		Specialization: "smart contract development, testing, software engineering, algorithmic trading, rockets",
		Department: "leadership",
		Desc: template.HTML(`
			Themis is a software engineer based in Athens. 
			He has significant experience developing software for a variety of projects such as web applications, algorithmic traders, and rockets. 
			He has worked on the Ethereum light client on Axelar and on a monitoring and alerting system for an EVM blockchain bridge. 
			He is pursuing a Master’s degree in Electrical and Computer Engineering from the National Technical University of Athens 
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
			Giulia is a doctoral student at the Technical University of Vienna (TU Wien), 
			where she <a href="https://dblp.org/pid/306/8308.html">focuses</a> on blockchain scalability, interoperability, and light clients
			She has been a summer research intern at Mysten Labs (2024) and at a16z crypto (2025). 
			Notably, in the <a href="https://eprint.iacr.org/2024/1680.pdf">Sunfish</a> paper she has formalized for the first time a new type of 
			blockchain client: the sparse client (aka. partially stateless client). 
			Among others, she has co-authored papers on light clients and bridges, including 
			<a href="https://eprint.iacr.org/2024/692.pdf">Blink</a>,
			<a href="https://eprint.iacr.org/2024/197.pdf">Alba</a>,
			<a href="https://eprint.iacr.org/2022/1721.pdf">Glimpse</a>.
		`),
		Image: "giulia_scaffino.png",
	},
	"nikos": {
		Handle:         "nikos",
		Name:           "Nikos Sfakianakis",
		Specialization: "software engineering, consensus, offensive security",
		Department: "engineering",
		Desc: template.HTML(`
		Nikos is a software engineer based in Athens. He is currently wrapping up his Master’s Degree in Electrical and Computer Engineering at the National Technical University of Athens
		where he is completing his thesis on Consensus Algorithms. 
		He has previously worked at Mysten Labs as a Blockchain Engineer and completed an internship at the Cybersecurity Department of OTE Group. 
		He has also competed in several Ethical Hacking and Offensive Security competitions.
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
		Department: "leadership",
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
		Nikola has worked as a full-stack web developer for ten years before joining Common Prefix. 
		During his career he worked on web applications, back-end services, WebAssembly and serverless applications, and has helped many startups grow. 
		At Common Prefix, Nikola is working on Bitcoin-related projects, smart contracts for Ethereum, CosmWasm and Solana, and cross-chain bridging. 
		He worked on various bridges including the Bitcoin Ordinals to EVM bridge, Filecoin IPC subnets on Bitcoin, and lead the integration of Hedera to Axelar's Amplifier bridging infrastructure.
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
	"sravya": {
		Handle:         "sravya",
		Name:           "Dr. Sravya Yandamuri",
		Specialization: "consensus, distributed systems, applied cryptography",
		Department: "science",
		Desc: template.HTML(`
			Sravya obtained her PhD  in Distributed Systems at Duke University under the guidance of Professor Kartik Nayak with a thesis on the topic of Byzantine Agreement in Asynchronous networks. 
			Previously, she was an intern at the now-defunct VMware Research as well as a visiting researcher at the CISPA Helmholtz Center for Information Security. 
			Her research includes topics such as State Machine Replication, Dynamic Availability, and Parallel Broadcast. 
			Her work has been published in notable conferences including CRYPTO and PODC.
			Highlights of her research include the papers
			<a href="https://eprint.iacr.org/2022/711.pdf">Efficient and Adaptively Secure Asynchronous Binary Agreement via Binding Crusader Agreement</a>,
			<a href="https://eprint.iacr.org/2025/1003.pdf">Low-Latency Dynamically Available Total Order Broadcast</a>,
			and <a href="https://eprint.iacr.org/2025/1016.pdf">Leader Election with Poly-logarithmic Communication Per Party</a>. 
			`),
		Image: "sravya_yandamuri.jpg",
		
	},
	"dionysis": {
		Handle:         "dionysis",
		Name:           "Dr. Dionysis Zindros",
		Specialization: "consensus, light clients, bridges, interoperability, fast bootstrapping, algorithms, software engineering",
		Department: "leadership",
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
