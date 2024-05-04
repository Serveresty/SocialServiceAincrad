import React, { useEffect, useState } from 'react';
import Cookies from 'universal-cookie';

const Chat = () => {
    const cookies = new Cookies();
    const [messages, setMessages] = useState([]);
    const [ws, setWs] = useState(null);
    const [message, setMessage] = useState('');
 
    useEffect(() => {
        const authToken = cookies.get('authToken');
        const searchParams = new URLSearchParams(window.location.search);
        const id = searchParams.get('id');

        const websocket = new WebSocket(`ws://localhost:8080/messages`);
 
        websocket.onopen = () => {
            websocket.send(JSON.stringify({authToken, id}));
        };
 
        websocket.onmessage = (evt) => {
            const message = (evt.data);
            setMessages((prevMessages) =>
                [...prevMessages, message]);
        };
 
        websocket.onclose = () => {
        };
 
        setWs(websocket);
 
        return () => {
            websocket.close();
        };
    }, []);
 
    const sendMessage = () => {
        if (ws) {
            ws.send(JSON.stringify({message}));
            setMessage('');
        }
    };
 
    const handleInputChange = (event) => {
        setMessage(event.target.value);
    };
 
    return (
        <div>
            <h1>
                Real-time Updates
            </h1>
            {messages.map((message, index) =>
                <p key={index}>{message}</p>)}
            <input type="text" value={message}
                onChange={handleInputChange} />
            <button onClick={sendMessage}>
                Send Message
            </button>
        </div>
    );
};
   
  export default Chat;
