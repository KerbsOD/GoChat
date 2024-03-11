import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput";
import { connect, sendMsg } from "./api";

export default function Chat({username}) {
    const [chatHistory, setChatHistory] = useState([]);

    useEffect(() => {
        document.title = "GoChat";

        const handleMessage = (msg) => {
            setChatHistory(prevChatHistory => [...prevChatHistory, msg]);
        };

        connect(handleMessage, {username});
    }); 

    const send = (event) => {
        // Si la tecla es enter entonces usa la funcion sendMsg para enviar el valor
        if (event.keyCode === 13) { // enter
            sendMsg(event.target.value);
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
