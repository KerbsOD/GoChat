import "./Message.scss";

export default function Message ({message}) {
    let temp = JSON.parse(message);
    const {type, statusmessage, sender, body} = temp

    const {username, content} = body

    let postMessage = (<div><strong>{sender}</strong>: {body} </div>)

    if (statusmessage === 1) {
        postMessage = (<div><strong>{sender} has joined the chat </strong></div>)
    }

    if (statusmessage === 3) {
        postMessage = (<div><strong>{sender} has left the chat </strong></div>)
    }

    return (
        <div className="Message">
            {postMessage}
        </div>
    )
}

