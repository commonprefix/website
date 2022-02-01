import Head from "next/head"
import Layout from "components/layout"
import Member from "components/teamMember"
import DownloadButton from "components/downloadButton"
import { Box, Link } from "@chakra-ui/react"

export default function Harmony() {
  const members = [
    {
      name: "Dr. Zeta Avarikioti",
      description: (
        <Box>
          Zeta is a post-doctoral blockchain researcher at IST Austria and TU Wien working with professor Krzysztof
          Pietrzak and based in Vienna. She graduated with a PhD from ETH Zürich and holds an engineering degree from
          the National Technical University of Athens. She specializes in distributed systems, scaling blockchains via
          sharding and channels, and the analysis of cryptoeconomic incentives. Among other venues, she has published in
          Financial Crypto and SODA. Highlights of her research include the papers{" "}
          <Link color="blue.500" href="https://arxiv.org/pdf/1905.11360.pdf">
            BRICK: Asynchronous Payment Channels
          </Link>
          ,{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2019/1092.pdf">
            Cerberus Channels: Incentivizing Watchtowers for Bitcoin
          </Link>
          , and{" "}
          <Link color="blue.500" href="https://arxiv.org/pdf/1910.10434.pdf">
            Divide and Scale: Formalization of Distributed Ledger Sharding Protocols
          </Link>
          .
        </Box>
      ),
    },
    /*
    {
      name: "Aleksis Brezas",
      description: (
        <Box>
          Aleksis is a senior software engineer based in Thessaloniki. With a dozen years of software engineering
          experience under his belt, he has worked with numerous programming languages and systems, with a particular
          taste for functional programming languages. He has worked on scaling web applications and in integrating
          heterogeneous systems and in large commercial web applications, and is the lead software architect behind{" "}
          <Link color="blue.500" href="https://geekbot.com/">
            geekbot
          </Link>{" "}
          and{" "}
          <Link color="blue.500" href="https://dailysteals.com/">
            dailysteals
          </Link>
          . His blockchain development experience includes smart contract development as well as cryptocurrency
          algorithmic trading. In Harmony, he is working closely with the scientists to implement surgical changes in
          the consensus layer pertaining to incentives, fees, and slashing.
        </Box>
      ),
    },
    */
    {
      name: "Dr. Pyrros Chaidos",
      description: (
        <Box>
          Pyrros is a post-doctoral blockchain researcher at the University of Athens. He holds a PhD from the
          University College of London, an MSc in Information Security from the University College of London, an MSc in
          Theoretical Computer Science from the University of Athens, and a BA in Mathematics from the University of
          Athens. He specializes in zero knowledge proof systems, with applications on proof-of-stake blockchains and
          voting in particular. Among other venues, he has published in EUROCRYPT, ACM CCS, PKC, ESORICS, and the
          Journal of Cryptology. Highlights of his research include the papers{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2021/916.pdf">
            Mithril: Stake-based Threshold Multisignatures
          </Link>
          ,{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2016/263.pdf">
            Efficient Zero-Knowledge Arguments for Arithmetic Circuits in the Discrete Log Setting
          </Link>
          , and{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2016/368.pdf">
            Foundations of Fully Dynamic Group Signatures
          </Link>
          .
        </Box>
      ),
    },
    {
      name: "Dr. Dimitris Karakosta",
      description: (
        <Box>
          Dimitris is a post-doctoral blockchain researcher at the University of Edinburgh where he works with professor
          Aggelos Kiayias. He holds a PhD from the University of Edinburgh and an Electrical and Computer Engineering
          degree from the National Technical University of Athens. He specializes in cryptocurrency wallets and
          macroeconomics, as well as chain checkpointing. Among other venues, he has published in Financial Crypto and
          presented at Black Hat Europe and Asia. Highlights of his research include the papers{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2019/034.pdf">
            A Formal Treatment of Hardware Wallets
          </Link>
          ,{" "}
          <Link color="blue.500" href="https://arxiv.org/pdf/1907.02434.pdf">
            Cryptocurrency Egalitarianism: A Quantitative Approach
          </Link>
          , and{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2020/173.pdf">
            Securing Proof-of-Work Ledgers via Checkpointing
          </Link>
          .
        </Box>
      ),
    },
    {
      name: "Prof. Dr. Lefteris Kokoris-Kogias",
      description: (
        <Box>
          Lefteris is an assistant professor at IST Austria based in Vienna. He holds a PhD from EPFL, and an Electrical
          and Computer Engineering degree from the National Technical University of Athens. He specializes in the
          systems aspects of blockchains with a focus on scalability, decentralized randomness generation, and software
          update dispersion. Among other venues, he has published in IEEE S&P (Oakland), USENIX Security, ACM CCS, PODC,
          ESORICS, and Financial Crypto. Highlights of his research include the papers{" "}
          <Link color="blue.500" href="https://infoscience.epfl.ch/record/255586/files/OmniLedger.pdf">
            Omniledger: A secure, scale-out, decentralized ledger via sharding
          </Link>
          ,{" "}
          <Link
            color="blue.500"
            href="https://scholar.google.com/citations?view_op=view_citation&hl=en&user=7TOIlqkAAAAJ&citation_for_view=7TOIlqkAAAAJ:kw52XkFRtyQC"
          >
            Enhancing bitcoin security and performance with strong consistency via collective signing
          </Link>
          , and{" "}
          <Link color="blue.500" href="https://discovery.ucl.ac.uk/id/eprint/10116632/1/Jovanovic_randomness.pdf">
            Scalable bias-resistant distributed randomness
          </Link>
          .
        </Box>
      ),
    },
    {
      name: "Dimitris Lamprinos",
      description: (
        <Box>
          Dimitris is a software engineer based in Thessaloniki. He holds a degree in Computer Science from Aristotle
          University of Thessaloniki. He works on smart contract development and basic consensus development. He has
          significant experience in developing and scaling web applications as well as blockchain development in the
          context of algorithmic trading. In the past, he has worked for{" "}
          <Link color="blue.500" href="https://www.amondo.com/">
            Amondo
          </Link>{" "}
          and{" "}
          <Link color="blue.500" href="https://geekbot.com/">
            Geekbot
          </Link>
          . In Harmony, he is working closely with the scientists to implement surgical changes in the consensus layer
          pertaining to incentives, fees, and slashing.
        </Box>
      ),
    },
    {
      name: "Dr. Dionysis Zindros",
      description: (
        <Box>
          Dionysis is a post-doctoral blockchain researcher at Stanford University. He holds a PhD from the University
          of Athens and an Electrical and Computer Engineering degree from the National Technical University of Athens.
          He specializes in blockchain scalability via light clients and fast bootstrapping. Among other venues he has
          published in IEEE S&P (Oakland), ACM CCS, ESORICS, and Financial Crypto, and presented at Black Hat Europe and
          Asia. Highlights of his research include the papers{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2017/963.pdf">
            Non-Interactive Proofs of Proof-of-Work
          </Link>
          ,{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2018/1239.pdf">
            Proof-of-Stake Sidechains
          </Link>
          , and{" "}
          <Link color="blue.500" href="https://eprint.iacr.org/2018/1048.pdf">
            Proof-of-Work Sidechains
          </Link>
          .
        </Box>
      ),
      image: "https://dionyziz.com/w/images/dionysis.png",
    },
  ]
  return (
    <Layout>
      <Head>
        <title>Common Prefix - Harmony</title>
      </Head>
      <span>
        Common Prefix is a team of scientists and engineers specializing in blockchains. We work from first principles,
        with an eye for rigor and formal cryptographic security proofs. We contribute as the advisory team for Harmony's
        underlying science. With a focus on consensus foundations and experts in the areas of proof-of-stake, auctions,
        layer 2, sharding, channels, bootstrapping, light clients, wallets, smart contracts, and interoperability, we
        review and recommend changes in all the fundamentals and protocol parameters. We deliver concrete
        recommendations for changes in the protocol and codebase, technical reports for dissemination within the
        community, and peer-reviewed research papers published in top-tier academic conferences describing our findings.
      </span>
      <DownloadButton file="/harmony_tech_report.pdf" />
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
