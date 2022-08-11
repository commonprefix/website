import Head from "next/head"
import Layout from "components/layout"
import { Image } from "@chakra-ui/react"
import Member from "components/teamMember"
import DownloadButton from "components/downloadButton"
import { Box, Link } from "@chakra-ui/react"
import team from "../team"

export default function FLRFinance() {
  const members = [
    team.zeta,
    team.aleksis,
    team.dimitris,
    team.stavros,
    team.pkakelas,
    team.dionysis,
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
        <Image minWidth="190" width="190" height="190" src="/cp_logo.jpg" m="2" />
        <Image height="75" src="/flr_finance_logo.png" m="2" />
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
