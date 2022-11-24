import {
  Box,
  TableContainer,
  Table,
  TableCaption,
  Thead,
  Tr,
  Th,
  Tbody,
  Td,
  Button,
  Popover,
  PopoverTrigger,
  Stack,
  Input,
  ButtonGroup,
  useDisclosure,
  PopoverContent,
  PopoverArrow,
  PopoverCloseButton,
  Text,
  Select,
  Flex,
  Center,
  FormControl,
  FormLabel,
} from "@chakra-ui/react";
import React, { useCallback, useEffect, useRef, useState } from "react";
import BigDecimal from "js-big-decimal";
import { ethers } from "ethers";
import { ABI, address } from "../contracts/PaymentContract";

const FrontPage = () => {
  const [receivers, setReceivers] = useState([]);

  const newReceiverAddressRef = useRef(null);
  const newReceiverValueRef = useRef(null);

  const [currency, setCurrency] = useState("ETH");

  const { onOpen, onClose, isOpen } = useDisclosure();
  const firstFieldRef = React.useRef(null);

  const handleClick = () => {
    setReceivers((receivers) => [
      ...receivers,
      {
        address: newReceiverAddressRef.current.value,
        value: newReceiverValueRef.current.value,
      },
    ]);
    onClose();
  };


  const handleCurrencyChange = (event) => {
    setCurrency(event.target.value);
  };

  const getReceiversTotal = () => {
    let total = new BigDecimal(0);

    receivers.forEach((receiver) => {
      total = total.add(new BigDecimal(receiver.value));
    });

    return total;
  };

  const getDepositMinimum = () => {
    const receiversTotal = getReceiversTotal();
    const fee = getFee();

    return receiversTotal.add(fee);
  };

  const getFee = () => {
    let total = getReceiversTotal();
    return total.multiply(new BigDecimal("0.03"));
  };

  const FormAddReceiver = ({ firstFieldRef, onCancel }) => {
    return (
      <Stack spacing={4}>
        <Text>Address:</Text>
        <Input placeholder="0x0.." id="address" ref={newReceiverAddressRef} />
        <Text>Value:</Text>
        <Input placeholder="0" id="value" ref={newReceiverValueRef} />
        <ButtonGroup display="flex" justifyContent="flex-end">
          <Button variant="outline" onClick={onCancel}>
            Cancel
          </Button>
          <Button onClick={handleClick} colorScheme="teal" size="md">
            Save
          </Button>
        </ButtonGroup>
      </Stack>
    );
  };

  return (
    <Box maxW="4xl" mx={"auto"} pt={5} px={{ base: 2, sm: 12, md: 17 }}>
      <br></br>

      <Flex
        flexDirection="row"
        justifyContent="space-between"
        alignItems="center"
        width="100%"
        as="nav"
        p={4}
        mx="auto"
        maxWidth="1150px"
        marginBottom="40px"
        maxW="8xl"
      >
        <Box>
          <Popover
            isOpen={isOpen}
            initialFocusRef={firstFieldRef}
            onOpen={onOpen}
            onClose={onClose}
            placement="right"
            closeOnBlur={false}
          >
            <PopoverTrigger>
              <Button colorScheme="teal" size="md">
                Add Receiver
              </Button>
            </PopoverTrigger>
            <PopoverContent p={5}>
              <PopoverArrow />
              <PopoverCloseButton />
              <FormAddReceiver
                firstFieldRef={firstFieldRef}
                onCancel={onClose}
              />
            </PopoverContent>
          </Popover>
        </Box>
        <Box>
          <Select onClick={handleCurrencyChange}>
            <option value="ETH">ETH</option>
            <option value="DAI">DAI</option>
            <option value="USDC">USDC</option>
          </Select>
        </Box>
      </Flex>

      {receivers && (
        <TableContainer>
          <Table variant="simple">
            <TableCaption>Receivers</TableCaption>
            <Thead>
              <Tr>
                <Th>Number</Th>
                <Th>Address</Th>
                <Th isNumeric>Value </Th>
              </Tr>
            </Thead>
            <Tbody>
              {receivers.map((receiver, index) => (
                <Tr>
                  <Td>{index + 1}</Td>
                  <Td>{receiver.address}</Td>
                  <Td isNumeric>
                    {receiver.value} {currency}
                  </Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </TableContainer>
      )}

      <Text>Receivers Total: {getReceiversTotal().getPrettyValue()}</Text>
      <Text>Fee: {getFee().getPrettyValue()}</Text>
      <br></br>
      <Text>Deposit Minimum: {getDepositMinimum().getPrettyValue()}</Text>

      <Box maxW="2xl" mx={"auto"} pt={10} px={{ base: 20, sm: 15, md: 20 }}>
        <FormLabel color="white">Deposit Value</FormLabel>
        <Input
          placeholder={getDepositMinimum().getPrettyValue()}
          // onChange={(event) => setAddress(event.target.value)}
          type="text"
          // value={address}
        />
        <br></br>
        <br></br>
        <Button
          marginTop={4}
          type="submit"
          colorScheme="teal"
        >
          Deposit
        </Button>
      </Box>

      <Box></Box>
    </Box>
  );
};

export default FrontPage;
