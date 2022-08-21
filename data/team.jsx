import { Box } from "@chakra-ui/react"
import Link from "../components/link"

const members = {
  shresth: {
    name: "Shresth Agrawal",
    description: (
      <Box>
        Shresth is a smart contract and backend developer.
        He holds a bachelor degree in Computer Science from Jacobs University Bremen.
        He has experience building efficient and secure algorithms,
        protocols, and smart contracts for several DeFi protocols.
        Previously he worked at
        {" "}
        <Link href="https://www.paraswap.io/">
          ParaSwap
        </Link>
        {" "}
        where he was responsible for developing a significant part of the core aggregation algorithm.
        He is interested in Cryptography, Security, Consensus Protocols, Decentralised Finance, and Ethereum.
      </Box>
    ),
    image: "team/shresth_agrawal.jpg",
  },
  zeta: {
    name: "Dr. Zeta Avarikioti",
    description: (
      <Box>
        Zeta is an upcoming Assistant Professor at the Technical University of Vienna (TU Wien) in Austria.
        She is currently a post-doctoral blockchain researcher at TU Wien working with professor Matteo Maffei.
        She graduated with a PhD from ETH Zürich and holds an engineering degree from the
        National Technical University of Athens and a masters degree from National Kapodistian University of Athens.
        She specializes in distributed systems, scaling blockchains via sharding and channels,
        and the analysis of cryptoeconomic incentives.
        Among other venues, she has published in Financial Crypto, AFT and SODA.
        Highlights of her research include the papers{" "}
        <Link href="https://arxiv.org/pdf/1905.11360.pdf">
          BRICK: Asynchronous Payment Channels
        </Link>
        ,{" "}
        <Link href="https://eprint.iacr.org/2019/1092.pdf">
          Cerberus Channels: Incentivizing Watchtowers for Bitcoin
        </Link>
        , and{" "}
        <Link href="https://arxiv.org/pdf/1910.10434.pdf">
          Divide and Scale: Formalization of Distributed Ledger Sharding Protocols
        </Link>
        .
      </Box>
    ),
    image: "team/zeta_avarikioti.jpeg",
  },
  aleksis: {
    name: "Aleksis Brezas",
    description: (
      <Box>
        Aleksis is a senior software engineer based in Thessaloniki. With a dozen years of software engineering
        experience under his belt, he has worked with numerous programming languages and systems, with a particular
        taste for functional programming languages. He has worked on scaling web applications and in integrating
        heterogeneous systems and in large commercial web applications, and is the lead software architect behind{" "}
        <Link href="https://geekbot.com/">
          geekbot
        </Link>{" "}
        and{" "}
        <Link href="https://dailysteals.com/">
          dailysteals
        </Link>
        . His blockchain development experience includes smart contract development as well as cryptocurrency
        algorithmic trading.
      </Box>
    ),
    image: "team/aleksis_brezas.jpg",
  },
  pyrros: {
    name: "Dr. Pyrros Chaidos",
    description: (
      <Box>
        Pyrros is a post-doctoral blockchain researcher at the University of Athens. He holds a PhD from the
        University College of London, an MSc in Information Security from the University College of London, an MSc in
        Theoretical Computer Science from the University of Athens, and a BA in Mathematics from the University of
        Athens. He specializes in zero knowledge proof systems, with applications on proof-of-stake blockchains and
        voting in particular. Among other venues, he has published in EUROCRYPT, ACM CCS, PKC, ESORICS, and the
        Journal of Cryptology. Highlights of his research include the papers{" "}
        <Link href="https://eprint.iacr.org/2021/916.pdf">
          Mithril: Stake-based Threshold Multisignatures
        </Link>
        ,{" "}
        <Link href="https://eprint.iacr.org/2016/263.pdf">
          Efficient Zero-Knowledge Arguments for Arithmetic Circuits in the Discrete Log Setting
        </Link>
        , and{" "}
        <Link href="https://eprint.iacr.org/2016/368.pdf">
          Foundations of Fully Dynamic Group Signatures
        </Link>
        .
      </Box>
    ),
    image: "team/pyrros_chaidos.jpeg",
  },
  nikolas: {
    name: "Nikolaos Kamarinakis",
    description: (
      <Box>
        Nikolas is a software engineer based in Athens, Greece and Maryland, USA.
        He is currently pursuing a bachelor's degree in Computer Science at the University of Maryland,
        specializing in Cybersecurity.
        He has multiple years of experience in full-stack software engineering and open-source development,
        as well as some expierience in offensive security.
        Nikolas is currently focused on smart contract development at Common Prefix
        as well as at the
        {" "}
        <Link href="https://researchdao.io">
          Research DAO
        </Link>,
        {" "}
        where he also coordinates operations.
      </Box>
    ),
    image: "/team/nikolaos_kamarinakis.jpg",
  },
  dimitris: {
    name: "Dr. Dimitris Karakostas",
    description: (
      <Box>
        Dimitris is a post-doctoral blockchain researcher at the University of Edinburgh where he works with professor
        Aggelos Kiayias. He holds a PhD from the University of Edinburgh and an Electrical and Computer Engineering
        degree from the National Technical University of Athens. He specializes in cryptocurrency wallets and
        macroeconomics, as well as chain checkpointing. Among other venues, he has published in Financial Crypto and
        presented at Black Hat Europe and Asia. Highlights of his research include the papers{" "}
        <Link href="https://eprint.iacr.org/2019/034.pdf">
          A Formal Treatment of Hardware Wallets
        </Link>
        ,{" "}
        <Link href="https://arxiv.org/pdf/1907.02434.pdf">
          Cryptocurrency Egalitarianism: A Quantitative Approach
        </Link>
        , and{" "}
        <Link href="https://eprint.iacr.org/2020/173.pdf">
          Securing Proof-of-Work Ledgers via Checkpointing
        </Link>
        .
      </Box>
    ),
    image: "/team/dimitris_karakostas.jpg",
  },
  lefteris: {
    name: "Prof. Dr. Lefteris Kokoris-Kogias",
    description: (
      <Box>
        Lefteris is an assistant professor at IST Austria based in Vienna. He holds a PhD from EPFL, and an Electrical
        and Computer Engineering degree from the National Technical University of Athens. He specializes in the
        systems aspects of blockchains with a focus on scalability, decentralized randomness generation, and software
        update dispersion. Among other venues, he has published in IEEE S&P (Oakland), USENIX Security, ACM CCS, PODC,
        ESORICS, and Financial Crypto. Highlights of his research include the papers{" "}
        <Link href="https://infoscience.epfl.ch/record/255586/files/OmniLedger.pdf">
          Omniledger: A secure, scale-out, decentralized ledger via sharding
        </Link>
        ,{" "}
        <Link
          href="https://scholar.google.com/citations?view_op=view_citation&hl=en&user=7TOIlqkAAAAJ&citation_for_view=7TOIlqkAAAAJ:kw52XkFRtyQC"
        >
          Enhancing bitcoin security and performance with strong consistency via collective signing
        </Link>
        , and{" "}
        <Link href="https://discovery.ucl.ac.uk/id/eprint/10116632/1/Jovanovic_randomness.pdf">
          Scalable bias-resistant distributed randomness
        </Link>
        .
      </Box>
    ),
    image: "/team/lefteris_kokoris-kogias.jpeg"
  },
  stavros: {
    name: "Stavros Korokithakis",
    description: (
      <Box>
        Stavros is a software developer based in Thessaloniki, Greece. He holds a bachelor's degree in Computer Science
        and Electronics Engineering from the Technical University of Crete. He has been programming for twenty-five years,
        with a focus on designing, developing, scaling, and maintaining web backends as well as on systems and data security.
        Stavros has worked in various sectors such as security, developing
        {" "}
        <Link href="https://www.silentcircle.com/products-and-solutions/silent-phone">
          Silent Phone
        </Link>
        ,{" "}
        alongside Phil Zimmermann and Jon Callas, in healthcare, and in blockchain, developing the Internet Computer at DFINITY.
      </Box>
    ),
    image: "/team/stavros_korokithakis.jpeg",
  },
  orfeas: {
    name: "Dr. Orfeas Stefanos Thyfronitis Litos",
    description: (
      <Box>
        Orfeas obtained his{" "}
        <Link href="https://gitlab.com/orfeasLitos/thesis/-/raw/master/thesis.pdf?inline=false">
          PhD in Cryptography and Blockchains
        </Link>{" "}
        at the University of Edinburgh in 2021. He has worked on{" "}
        <Link href="http://fc17.ifca.ai/preproceedings/paper_37.pdf">
          building and analyzing
        </Link>{" "}
        decentralized applications on blockchains,{" "}
        <Link href="https://fc20.ifca.ai/preproceedings/132.pdf">
          layer-2 protocols
        </Link>,{" "}
        <Link href="https://eprint.iacr.org/2021/747.pdf">
          payment channels
        </Link>, and{" "}
        <Link href="https://tokenomics2019.org/Documents/proceedings-full.pdf#page=17">
          voting solutions
        </Link>,{" "}
        all through formal cryptographic methods. Among others, he has{" "}
        <Link href="https://ieeexplore.ieee.org/abstract/document/9155145">
          formally analyzed the security of the Lightning Network
        </Link>{" "}
        and created a{" "}
        <Link href="https://eprint.iacr.org/2021/747.pdf">
          novel virtual payment channel construction
        </Link>.{" "}
        He is knowledgeable in{" "}
        <Link href="https://github.com/OrfeasLitos/virtual-payment-channels/tree/master/simulation">
          software engineering
        </Link>{" "}
        and{" "}
        <Link href="https://github.com/OrfeasLitos/cerberus-script">
          secure architecture
        </Link>.{" "}
        His interests further include formal verification,{" "}
        <Link href="http://fc22.ifca.ai/preproceedings/69.pdf">
          incentive analysis
        </Link>,{" "}
        and provable security.
      </Box>
    ),
    image: "/team/orfeas_litos.jpg",
  },
  pkakelas: {
    name: "Dimitris Lamprinos",
    description: (
      <Box>
        Dimitris is a software engineer based in Thessaloniki. He holds a Bachelor's degree in Computer Science from Aristotle
        University of Thessaloniki. He works on smart contract development and basic consensus development. He has
        significant experience in developing and scaling web applications as well as blockchain development in the
        context of algorithmic trading. In the past, he has worked for{" "}
        <Link href="https://www.amondo.com/">
          Amondo
        </Link>{" "}
        and{" "}
        <Link href="https://geekbot.com/">
          Geekbot
        </Link>
        .
      </Box>
    ),
    image: "/team/dimitris_lamprinos.jpeg",
  },
  themis: {
    name: "Themis Papameletiou",
    description: (
      <Box>
        Themis is a software engineer based in Athens.
        He has significant experience developing software for a variety of projects
        such as web applications, algorithmic traders and rockets.
        He is pursuing a Master's degree in Electrical and Computer Engineering
        from the National Technical University of Athens
        and has also completed internships at Google and at the European Space Agency.
        Currently, he is focused on smart contract development and testing.
      </Box>
    ),
    image: "/team/themis_papameletiou.jpg",
  },
  apostolos: {
    name: "Apostolos Tzinas",
    description: (
      <Box>
        Apostolos is a smart contract software engineer.
        He is pursuing a joint Bachelor and Master's degree in Electrical and Computer Engineering
        at the National Technical University of Athens.
        Apostolos has extensive front-end software engineering experience working at
        {" "}
        <Link href="https://www.mayainsights.com/">
          Maya Insights
        </Link>
        ,{" "}
        as well as at NutriDice, where he took on a full-stack engineering role.
        He has worked with numerous programming languages and technical stacks.
      </Box>
    ),
    image: "team/apostolos_tzinas.jpg",
  },
  dionysis: {
    name: "Dr. Dionysis Zindros",
    description: (
      <Box>
        Dionysis is a post-doctoral blockchain researcher at Stanford University. He holds a PhD from the University
        of Athens and an Electrical and Computer Engineering degree from the National Technical University of Athens.
        He specializes in blockchain scalability via light clients and fast bootstrapping. Among other venues he has
        published in IEEE S&P (Oakland), ACM CCS, ESORICS, and Financial Crypto, and presented at Black Hat Europe and
        Asia. Highlights of his research include the papers{" "}
        <Link href="https://eprint.iacr.org/2017/963.pdf">
          Non-Interactive Proofs of Proof-of-Work
        </Link>
        ,{" "}
        <Link href="https://eprint.iacr.org/2018/1239.pdf">
          Proof-of-Stake Sidechains
        </Link>
        , and{" "}
        <Link href="https://eprint.iacr.org/2018/1048.pdf">
          Proof-of-Work Sidechains
        </Link>
        .
      </Box>
    ),
    image: "team/dionysis_zindros.png",
  },
}

module.exports = members