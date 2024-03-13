import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput";

export default function Chat({username}) {
    const [chatHistory, setChatHistory] = useState([]);
    const [socket, setSocket] = useState(null);

    useEffect(() => {
        document.title = "GoChat";

        const socket = new WebSocket("ws://localhost:8080/ws");
        
        console.log("Attempting Connection...");

        socket.onopen = () => {
            console.log("Successfully Connected!");
        };

        socket.onmessage = event => {
            let msg = event.data;

            console.log('Message recieved: ', msg);

            let temp = JSON.parse(msg);
            const {type, statusmessage, sender, body} = temp

            console.log("Debug")
            console.log(temp)

            if ( statusmessage === 0 ) {
                const message = {
                    username: username,
                    content: "Username provided"
                };
            
                sendMessage(message);
            }

            callback(event)
        };

        socket.onerror = error => {
            console.log("Socket Error: ", error);
        };

        socket.onclose = event => {
            console.log("Socket Closed Connection: ", event);
        };

        function sendMessage(message) {
            console.log("Sending Message to backend: ", message);
            socket.send(JSON.stringify(message))
        };

        setSocket({
            socket,
            sendMessage
        });

        const callback = (msg) => {
            setChatHistory(prevChatHistory => [...prevChatHistory, msg]);
        };

        return() => {
            socket.close()
        };
        
    }, []); 

    const send = (event) => {
        // Si la tecla es enter entonces usa la funcion sendMsg para enviar el valor
        if (event.keyCode === 13) { // enter
            const message = {
                username: username,
                content: event.target.value,
            };
            
            if (socket) {
                socket.sendMessage(message)
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
