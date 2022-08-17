import Head from "next/head"
import styles from "./styles/layout.module.scss"

export default function Layout({ title, icon, children }) {
  return (
    <>
      <Head>
        <meta charSet="UTF-8" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>{`Common Prefix${title ? " - " + title : ""}`}</title>
        <link rel="icon" type="image/png" href={icon || "/philosopher-stone.png"} />
      </Head>
      <div className={styles.container}>{children}</div>
    </>
  )
}
