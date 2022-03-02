import { Box, Image, Flex } from "@chakra-ui/react"

export default function Member({ member }) {
  return (
    <Flex mt="10" width="100%" flexWrap="wrap" borderWidth="1px" borderRadius="lg" overflow="hidden" alignItems="center" justifyContent="center">
      <Image m="6"borderRadius="full" boxSize="200px" src={member.image ? member.image : "/avatar.png"} alt={member.name} />
      <Box m="6" minWidth="10rem" maxWidth="35rem">
        <Box fontWeight="semibold" lineHeight="tight" isTruncated>
          {member.name}
        </Box>
        <Box display="flex" mt="2" color="gray.600" fontSize="sm">
          {member.description}
        </Box>
      </Box>
    </Flex>
  )
}
