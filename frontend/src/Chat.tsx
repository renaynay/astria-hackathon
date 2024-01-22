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

  const connect = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (username.trim() === '') {
      // Show a toast message to inform the user to enter a username
      toast({
        title: "Error",
        description: "Username cannot be empty.",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
      return;
    }
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
  
  const sendMessage = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (message.trim() === '') {
      // Show a toast message to inform the user to enter a message
      toast({
        title: "Error",
        description: "Message cannot be empty.",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
      return;
    }
    if (ws) {
      ws.send(message);
      setMessage('');
    }
  };

  return (
    <VStack spacing={4} align="stretch" m={8}>
      <Heading as="h1" size="lg" textAlign="center">Modular Chat</Heading>
      <form onSubmit={connect}>
          <Input
            placeholder="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <Button colorScheme="blue" type="submit" mt={2}>Join Chat</Button>
      </form>
      <Textarea
        placeholder="Messages"
        value={messages.join('\n')}
        readOnly
        height="300px"
      />
      <form onSubmit={sendMessage}>
        <Box>
          <Input
            placeholder="chat"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            mr={2}
          />
          <Button colorScheme="green" type="submit" mt={2}>Send</Button>
        </Box>
      </form>
    </VStack>
  );
};

export default Chat;