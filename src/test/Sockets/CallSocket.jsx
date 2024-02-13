import { Button, Input } from 'antd';
import React, { useEffect, useState } from 'react';
import useWebSocket from 'react-use-websocket';

export default function CallSocket() {
  const WS_URL = 'ws://localhost:8080/manager/ws?room=thang2';

  const [sendMessage, setSendMessage] = useState('');
  const [messages, setMessages] = useState([]);
  const [isSending, setIsSending] = useState(false);


  const userName = localStorage.getItem('user_name');

  console.log(userName)
  const {
    sendJsonMessage,
    lastJsonMessage,
    readyState,
  } = useWebSocket(WS_URL);

  useEffect(() => {
    if (lastJsonMessage && lastJsonMessage.user_name !== userName) {
      setMessages((prevMessages) => [...prevMessages, lastJsonMessage]);
    }
  }, [lastJsonMessage]);

  const handleSend = () => {
    if (!isSending) {
      setIsSending(true);
      const message = {
        content: sendMessage,
        user_name: userName,
      };
      sendJsonMessage(message);
      setSendMessage('');
      setIsSending(false);
    }
  };

  return (
    <div>
      <ul style={{ display: 'flex' }}>
        <li>
          <Input value={sendMessage} onChange={(e) => setSendMessage(e.target.value)} />
        </li>
        <li>
          <Button onClick={handleSend}>Send</Button>
        </li>
        <ul>
          {messages.map((message, index) => (
              <li key={index} style={{ color: message.user_name}}>
              {message.content},---,
              {/* {userName} */}
            </li>
          ))}
        </ul>
      </ul>
    </div>
  );
}