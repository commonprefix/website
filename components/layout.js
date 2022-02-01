import styles from "./styles/layout.module.scss"

export default function Layout({ children }) {
  return <div className={styles.container}>{children}</div>
}
