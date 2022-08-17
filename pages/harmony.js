import Layout from "components/layout"
import { Image } from '@chakra-ui/react'
import Member from "components/teamMember"
import DownloadButton from "components/downloadButton"
import { Box, Link } from "@chakra-ui/react"
import team from "../data/team"

export default function Harmony() {
  const members = [
    team.zeta,
    team.pyrros,
    team.dimitris,
    team.lefteris,
    team.pkakelas,
    team.dionysis,
  ]
  return (
    <Layout title="Harmony">
      <Box display="flex" flexWrap="wrap" alignItems="center" justifyContent="center" mb="10" >
        <Image minWidth="190" width="190" height="190" src="/cp_logo.jpg" />
        <Image minWidth="190" width="190" height="190" src="/harmony_logo.png" />
      </Box>
      <span>
        Common Prefix is a team of scientists and engineers specializing in blockchains. We work from first principles,
        with an eye for rigor and formal cryptographic security proofs. We contribute as the advisory team for Harmony's
        underlying science. With a focus on consensus foundations and experts in the areas of proof-of-stake, auctions,
        layer 2, sharding, channels, bootstrapping, light clients, wallets, smart contracts, and interoperability, we
        review and recommend changes in all the fundamentals and protocol parameters. We deliver concrete
        recommendations for changes in the protocol and codebase, technical reports for dissemination within the
        community, and peer-reviewed research papers published in top-tier academic conferences describing our findings.
      </span>
      <DownloadButton file="/harmony_privacy_report.pdf" text="DOWNLOAD PRIVACY REPORT"/>
      <DownloadButton file="/harmony_tech_report_v2.pdf" text="DOWNLOAD TECH REPORT" version="2"/>
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
