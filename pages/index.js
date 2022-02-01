import Head from 'next/head'
import Image from 'next/image'
import Layout from 'components/layout'

export default function Home() {
  return (
    <Layout>
      <Head>
        <meta charSet="UTF-8"/>
        <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
        <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
        <title>Common Prefix</title>

        <link href="https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css" rel="stylesheet"/>
        <link rel="stylesheet" type="text/css" href="https://cdn.rawgit.com/dreampulse/computer-modern-web-font/master/fonts.css"/>
        <link rel="icon" type="image/png" href="/philosopher-stone.png"/>
      </Head>
      <Image width="190" height="190" src="/transmutation-circle.jpg"/>
      <h1 className='mx-auto m-4 text-center text-xl font-bold'>Common Prefix</h1>

      <div className='text-justify p-2 max-w-sm m-auto'>
        <p className='my-4'>We are a small team of scientists and software engineers
        offering blockchain science consulting services.</p>

        <p className='my-4'>We work using rigorous cryptographic techniques to
        design simple and provably secure protocols from first
        principles and guide developers in implementing them
        successfully.</p>

        <p className='my-4'>For a quote, reach out to{" "}
        <a className='text-pink-600 hover:underline' href='mailto:hello@commonprefix.com'>hello@commonprefix.com</a>.</p>
      </div>
    </Layout>
  )
}
