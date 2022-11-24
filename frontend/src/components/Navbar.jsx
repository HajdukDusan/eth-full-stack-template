import {
  Button,
  Flex,
  Box,
  Text,
  useColorMode,
  Center,
} from "@chakra-ui/react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <>
      <Box maxW="2xl" mx={"auto"} pt={1} px={{ base: 2, sm: 12, md: 17 }}>
        <Center>
          <Text fontSize="6xl">Title</Text>
        </Center>
      </Box>
    </>
  );
};

export default Navbar;
