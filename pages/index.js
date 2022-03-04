import Head from "next/head"
import Image from "next/image"
import Layout from "components/layout"
import { Box, Link } from "@chakra-ui/react"

export default function Home() {
  return (
    <Layout>
      <Head>
        <meta charSet="UTF-8" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Common Prefix</title>

        <link rel="icon" type="image/png" href="/philosopher-stone.png" />
      </Head>
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
          <Link color="pink.500"  href="mailto:hello@commonprefix.com">
            hello@commonprefix.com
          </Link>
          .
        </p>
      </Box>
    </Layout>
  )
}
