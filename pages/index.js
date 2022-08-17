import Image from "next/image"
import Layout from "components/layout"
import { Box } from "@chakra-ui/react"
import Link from "../components/link"

export default function Home() {
  return (
    <Layout>
      <Image width="190" height="190" src="/cp_logo.jpg" />
      <Box mt="5" fontSize="23" color="black"><h1>Common Prefix</h1></Box>

      <Box mt="10" maxWidth="25rem">
        <p>
          We are a small team of scientists and software engineers offering blockchain science consulting services.
        </p>
        <p>
          We work using rigorous cryptographic techniques to design simple and provably secure protocols from first
          principles and guide developers in implementing them successfully.
        </p>
        <p>
          For a quote, reach out to{" "}
          <Link href="mailto:hello@commonprefix.com">
            hello@commonprefix.com
          </Link>
          .
        </p>
      </Box>
    </Layout>
  )
}
