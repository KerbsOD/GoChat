import "./Message.scss";

export default function Message ({message}) {
    let temp = JSON.parse(message);
    const { sender, body } = temp
    
    return (
        <div className="Message">
            <div><strong>{sender}</strong>: {body}</div>
        </div>
    )
}

