import { Button, Input, message } from 'antd'
import React, { useEffect, useState } from 'react'
import useWebSocket, { ReadyState } from "react-use-websocket"

export default function CallSocket() {

    const WS_URL = 'ws://localhost:8080/manager/ws'

    const [sendMessage, setSendMessage] = useState('');
    const [receivedMessage, setReceivedMessage] = useState('');

    const {
        sendJsonMessage,
        lastJsonMessage,
        readyState,
    } = useWebSocket(WS_URL);

    useEffect(() => {
        if (lastJsonMessage) {
            setReceivedMessage(JSON.stringify(lastJsonMessage));
        }
    }, [lastJsonMessage]);

    const handleSend = () => {

        const message = {
            text: sendMessage,
            user_name: 'Your Name', // Thay 'Your Name' bằng tên của bạn
            timestamp: new Date().toISOString(),
          };

        sendJsonMessage(message);

        setSendMessage('');
    };

    return (
        <div>
            <ul style={{ display: 'flex' }}>
                <li>
                    <Input
                        value={sendMessage}
                        onChange={(e) => setSendMessage(e.target.value)}
                    />                </li>
                <li>
                    <Button onClick={handleSend}>Send</Button>
                </li>
                <ul>
                    <li style={{color:'red'}}>{sendMessage}</li>
                    <li style={{color:'blue'}}>{receivedMessage}</li>
                </ul>
            </ul>

        </div>
    )
}
