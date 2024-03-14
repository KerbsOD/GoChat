import React, { useEffect, useRef } from 'react';
import "./ChatHistory.scss";
import Message from '../Message/'

export default function ChatHistory({ chatHistory }) {
    const chatHistoryRef = useRef(null);

    const scrollToBottom = () => {
        if (chatHistoryRef.current) {
            chatHistoryRef.current.scrollTop = chatHistoryRef.current.scrollHeight;
        }
    };

    useEffect(() => {
        scrollToBottom();
    }, [chatHistory]);

    return (
        <div ref={chatHistoryRef} className='ChatHistory'>
            {chatHistory.map((msg, index) => (
                <div key={index} className='MessageContainer'>
                    <Message message={msg} />
                </div>
            ))}
        </div>
    );
}
