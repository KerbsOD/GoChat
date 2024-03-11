import "./Message.scss";

export default function Message ({message}) {
    let temp = JSON.parse(message);
    const {statusmessage, sender, body,} = temp

    let postMessage = (<div><strong>{sender}</strong>: {body}</div>)

    if (statusmessage === 0) {
        postMessage = (<div><strong>{sender} has joined the chat </strong></div>)
    }

    if (statusmessage === 2) {
        postMessage = (<div><strong>{sender} has left the chat </strong></div>)
    }

    return (
        <div className="Message">
            {postMessage}
        </div>
    )
}

