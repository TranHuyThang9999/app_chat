import React, { useEffect, useState } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';

const WS_URL = 'ws://localhost:8050/ws';

export default function AppSocket() {
  const [socketUrl, setSocketUrl] = useState(WS_URL);
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState([]);
  const [userName, setUserName] = useState('');

  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

  useEffect(() => {
    if (lastMessage !== null) {
      const receivedMessage = JSON.parse(lastMessage.data);
      setMessages((prevMessages) => [...prevMessages, receivedMessage]);
    }
  }, [lastMessage]);

  const handleSendMessage = () => {
    if (message.trim() !== '' && userName.trim() !== '') {
      const newMessage = {
        text: message.trim(),
        userName: userName.trim(),
      };
      sendMessage(JSON.stringify(newMessage));
      setMessage('');
    }
  };

  return (
    <div>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
      />
      <input
        type="text"
        value={userName}
        onChange={(e) => setUserName(e.target.value)}
      />
      <button onClick={handleSendMessage}>Send</button>
      {readyState === ReadyState.OPEN ? (
        <span>Connected</span>
      ) : (
        <span>Connecting...</span>
      )}
      <ul style={{ color: 'red' }}>
        {messages.map((msg, index) => (
          <li key={index}>
            {msg.text} - {new Date(msg.timestamp).toLocaleString()} - {msg.userName}
          </li>
        ))}
      </ul>
    </div>
  );
}