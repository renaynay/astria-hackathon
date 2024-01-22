import { useState } from 'react';
import {
  Box,
  Button,
  Input,
  VStack,
  Heading,
  Textarea,
  useToast,
} from '@chakra-ui/react';

const Chat = () => {
  const [username, setUsername] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const [messages, setMessages] = useState<string[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);
  const toast = useToast();

  const connect = () => {
    const websocket = new WebSocket("ws://localhost:3000/websocket");

    websocket.onopen = () => {
      toast({
        title: "Connected",
        description: "You are now connected to the chat server.",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
      websocket.send(username);
    };

    websocket.onclose = () => {
      toast({
        title: "Disconnected",
        description: "You have been disconnected from the chat server.",
        status: "warning",
        duration: 5000,
        isClosable: true,
      });
    };

    websocket.onmessage = (e) => {
      setMessages(prevMessages => [...prevMessages, e.data]);
    };

    setWs(websocket);
  };

  const sendMessage = () => {
    if (ws) {
      ws.send(message);
      setMessage('');
    }
  };

  return (
    <VStack spacing={4} align="stretch" m={8}>
      <Heading as="h1" size="lg" textAlign="center">Modular Chat</Heading>
      <Input
        placeholder="username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <Button colorScheme="blue" onClick={connect}>Join Chat</Button>
      <Textarea
        placeholder="Messages"
        value={messages.join('\n')}
        readOnly
        height="300px"
      />
      <Box>
        <Input
          placeholder="chat"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          mr={2}
        />
        <Button colorScheme="green" onClick={sendMessage} mt={4}>Send</Button>
      </Box>
    </VStack>
  );
};

export default Chat;