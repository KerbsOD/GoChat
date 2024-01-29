import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message);
        this.state = {
            message: temp
        };
    }

    render() {
        const { sender, body } = this.state.message;
        return( 
            <div className="Message">
                <div><strong>{sender}</strong>: {body}</div>
            </div>
        );
    }
}

export default Message;
