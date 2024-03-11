import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput";
import "./App.css";
import { connect, sendMsg } from "./api";

export default function App() {
    const [chatHistory, setChatHistory] = useState([]);

    useEffect(() => {
        document.title = "GoChat";

        const handleMessage = (msg) => {
            console.log("New Message")
            setChatHistory(prevChatHistory => [...prevChatHistory, msg]);
            console.log(chatHistory);
        };

        connect(handleMessage);
    }); 

    const send = (event) => {
        if (event.keyCode === 13) { // enter
            sendMsg(event.target.value);
            event.target.value = "";
        }
    };

    return (
        <div className="App">
            <Header />
            <ChatHistory chatHistory={chatHistory} />
            <ChatInput send={send} />
        </div>
    );
};

