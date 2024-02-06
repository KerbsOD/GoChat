import "./ChatInput.scss";

export default function ChatInput({ send }) {
    return (
        <div className="ChatInput">
            <input onKeyDown={send}/>
        </div>
    )
}
    