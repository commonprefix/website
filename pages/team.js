import Layout from "components/layout"
import { Image, Heading  } from "@chakra-ui/react"
import Member from "components/teamMember"
import team from "../data/team"

export default function Team() {
  const members = Object.values(team)
  return (
    <Layout title="Team">
      <Image minWidth="190" width="190" height="190" src="/cp_logo.jpg" m={2} />
      <Heading mt={8}>Our Team</Heading>
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
