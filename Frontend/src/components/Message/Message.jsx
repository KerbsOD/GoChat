import "./Message.scss";

export default function Message ({ message }) {
    const { type, sender, body } = message;

    let postMessage = null;

    if (type === 1) {
        postMessage = <div><strong>{sender} has joined the chat</strong></div>;
    } else if (type === 3) {
        postMessage = <div><strong>{sender} has left the chat</strong></div>;
    } else {
        postMessage = (
            <div className="message-body">
                <div className="sender"><strong>{sender}</strong></div>
                <div className="body">{body}</div>
            </div>
        );
    }

    return (
        <div className="Message">
            {postMessage}
        </div>
    );
}

