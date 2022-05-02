import { Box, LinkBox, LinkOverlay } from "@chakra-ui/react"
import { DownloadIcon } from "@chakra-ui/icons"
import styles from "./styles/download.module.scss"

export default function DownloadButton({ file, text, version }) {
  return (
    <LinkBox className={styles.link_button} mt="10" href={file} download>
      <LinkOverlay href={file} download>
        <Box className={styles.button} display="flex" borderRadius="lg">
          <DownloadIcon />

          <Box ml="5" fontWeight="semibold" lineHeight="tight" isTruncated>
            {text} {version && "v"}{version}
          </Box>
        </Box>
      </LinkOverlay>
    </LinkBox>
  )
}
