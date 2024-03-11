import "./ChatInput.scss";

export default function ChatInput({ send }) {
    return (
        <div className="ChatInput">
            {/*Si apreto alguna tecla, usa la funcion send para handlearlo*/}
            <input onKeyDown={send}/>
        </div>
    )
}
    