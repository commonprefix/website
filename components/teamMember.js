import { Box, Image, Flex } from "@chakra-ui/react"

export default function Member({ member }) {
  const property = {
    imageUrl: "https://bit.ly/2Z4KKcF",
  }

  return (
    <Flex p="6" mt="10" borderWidth="1px" borderRadius="lg" overflow="hidden" alignItems="center">
      <Image borderRadius="full" boxSize="200px" src={member.image ? member.image : "/avatar.png"} alt={member.name} />

      <Box ml="10">
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
