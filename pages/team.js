import Head from "next/head"
import Layout from "components/layout"
import { Image, Heading  } from "@chakra-ui/react"
import Member from "components/teamMember"
import team from "../team"

export default function Team() {
  const members = Object.values(team)
  return (
    <Layout>
      <Head>
        <meta charSet="UTF-8" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Common Prefix - Team</title>
        <link rel="icon" type="image/png" href="/philosopher-stone.png" />
      </Head>
      <Image minWidth="190" width="190" height="190" src="/cp_logo.jpg" m={2} />
      <Heading mt={8}>Our Team</Heading>
      {members.map((member) => (
        <Member key={member.name} member={member} />
      ))}
    </Layout>
  )
}
