import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput";
import Header from './components/Header';

export default function Chat({username}) {
    const [chatHistory, setChatHistory] = useState([]);
    const [socket, setSocket] = useState(null);
    const navigate = useNavigate()
    
    
    useEffect(() => {
        document.title = "GoChat";

        if (username === '') {
            navigate('/')
            return
        }
    
        const socket = new WebSocket("ws://localhost:8080/ws");
        
        console.log("Attempting Connection...");

        socket.onopen = () => {
            console.log("Successfully Connected!");
        };

        socket.onmessage = event => {
            let message = JSON.parse(event.data);
            const {type, sender, body} = message
            console.log('Message recieved: ', message);
            
            if ( type === 0 ) {
                console.log("Username request received")
                sendMessage(username);
            } else {
                setChatHistory(prevChatHistory => [...prevChatHistory, message]);
            }
        };

        socket.onerror = error => {
            console.log("Socket Error: ", error);
        };

        socket.onclose = event => {
            console.log("Socket Closed Connection: ", event);
        };

        function sendMessage(message) {
            const messageJSON = {
                type: 2,
                sender: username,
                body: message,
            };
            console.log("Sending Message to backend: ", messageJSON);
            socket.send(JSON.stringify(messageJSON))
        };

        setSocket({
            socket,
            sendMessage
        });

        return() => {
            socket.close()
            console.log("Socket closed")
        };
        
    }, []); 

    const send = (event) => {
        if (event.key === 'Enter') { 
            if (socket) {
                socket.sendMessage(event.target.value)
            }
            event.target.value = "";
        }
    };

    return (
        <div className="Chat">
            <Header />
            <ChatHistory chatHistory={chatHistory} />
            <ChatInput send={send} />
        </div>
    );
};

