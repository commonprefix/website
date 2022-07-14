import Head from "next/head"
import Layout from "components/layout"
import { Image } from "@chakra-ui/react"
import Member from "components/teamMember"
import DownloadButton from "components/downloadButton"
import { Box, Link } from "@chakra-ui/react"

export default function FLRFinance() {
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
      image: "team/zeta_avarikioti.jpeg",
    },
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
    {
      name: "Dr. Dimitris Karakostas",
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
      image: "/team/dimitris_karakostas.jpg",
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
      image: "/team/dimitris_lamprinos.jpeg",
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
      image: "team/dionysis_zindros.png",
    },
  ]
  return (
    <Layout>
      <Head>
        <meta charSet="UTF-8" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Common Prefix - FLR Finance</title>
        <link rel="icon" type="image/png" href="/philosopher-stone.png" />
      </Head>
      <Box display="flex" flexWrap="wrap" alignItems="center" justifyContent="center" mb="10">
        <Image minWidth="190" width="190" height="190" src="/cp_logo.jpg" />
        <Image height="75" src="/flr_finance_logo.png" />
      </Box>
      <span>
        Common Prefix is a team of scientists and engineers specializing in blockchains. We analyze and design simple
        but rigorous protocols from first principles, with provable security in mind. Our consulting and audits pertain
        to theoretical cryptographic protocol analyses as well as the pragmatic auditing of implementations in both core
        consensus technologies and application layer smart contracts. In essence, we help ensure that FLR Finance can
        uphold its promise to keep its community safe by delivering concrete recommendations for improvements of its
        protocols and codebase as well as by enhancing its testing pipeline. Our technical reports are published openly
        for dissemination within the community. Additionally, we develop novel protocols that expand the FLR Finance
        ecosystem even further.
      </span>
      <DownloadButton file="/flr_finance_loans_post_mortem.pdf" text="LOANS INCIDENT POST-MORTEM" />
      <DownloadButton file="/flr_finance_ftso_reward_manager_audit.pdf" text="FTSORewardManager AUDIT" />
      <DownloadButton file="/flr_finance_price_feed_ftso_connector_audit.pdf" text="PriceFeedFtsoConnector AUDIT" />
      <DownloadButton file="/flr_finance_loans_audit.pdf" text="LoansStable AUDIT" />
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
