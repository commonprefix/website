import Head from "next/head"
import Layout from "components/layout"
import Member from "components/member"
import Download from "components/download"
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
      <Download />
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
