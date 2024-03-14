import "./ChatInput.scss";

export default function ChatInput({ send }) {
    return (
        <div className="ChatInput">
            <input id="chatInput" type="text" onKeyDown={send} placeholder="Type your message here..." />
        </div>
    )
}
    