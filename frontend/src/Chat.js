import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput";
import { connect, sendMsg } from "./api";

export default function Chat({username}) {
    const [chatHistory, setChatHistory] = useState([]);

    useEffect(() => {
        document.title = "GoChat";

        const callback = (msg) => {
            setChatHistory(prevChatHistory => [...prevChatHistory, msg]);
        };

        connect(callback, username);
    }, [username]); 

    const send = (event) => {
        // Si la tecla es enter entonces usa la funcion sendMsg para enviar el valor
        if (event.keyCode === 13) { // enter
            const message = {
                content: event.target.value,
                username: username,
            };
            
            const messageString = JSON.stringify(message)
            
            sendMsg(messageString);
            
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
