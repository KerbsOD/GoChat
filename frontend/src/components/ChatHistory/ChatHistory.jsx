import "./ChatHistory.scss";
import Message from '../Message/'

export default function ChatHistory({chatHistory}) {
    console.log(chatHistory)
    const messages = chatHistory.map(msg => <Message message={msg.data} />);
    return (
        <div className='ChatHistory'>
            {messages}
        </div>
    );
}


