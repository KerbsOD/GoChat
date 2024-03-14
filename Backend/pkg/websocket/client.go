package websocket

import (
    "encoding/json"
    "log"
	"github.com/gorilla/websocket"
)

// Conn is our websocket connection between the frontend and backend
// Clients send messages from the frontend to de back via Conn,
// then the message is sent to everyone on the pool via Pool.BroadcastMessage.
// At last each client sents the message received from the pool to the frontend via its own Conn to 
// update the UI.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Type represents what action is the client doing.
// 0: Requesting username 
// 1: Joined the pool 
// 2: Sent message to the pool 
// 3: Left the pool
type Message struct {
	Type          int    `json:"type"`
	Sender        string `json:"sender"`
	Body          string `json:"body"`
}

func NewClient(connection *websocket.Conn, chatPool *Pool) *Client {
    return &Client {
        ID: "Default",
        Conn: connection,
        Pool: chatPool,
    }
}

func (c *Client) RequestUsername() {
    if err := c.SendMessage(Message{Type: 0, Sender: "", Body: ""}); err != nil {
        log.Println("Error sending username request: ", err)
        return
    }
    
    message, err := c.ReadMessage()
    if err != nil {
        log.Println("Error reading username response: ", err)
        return
    }

    c.ID = message.Body
}

func (c *Client) SendMessage(message Message) error {
    if err := c.Conn.WriteJSON(message); err != nil {
        return err
    }
    return nil
}

func (c *Client) ReadMessage() (Message, error) {
    _, data, err := c.Conn.ReadMessage()
    if err != nil {
        return Message{}, err
    }

    var message Message
    if err := json.Unmarshal(data, &message); err != nil {
        return Message{}, err
    }
    
    return message, nil
}

func (c *Client) Listen() {
    defer func() {
        c.Pool.Unregister <- c 
        c.Conn.Close()
    }()

    for {
        message, err := c.ReadMessage()
        if err != nil {
            log.Println("Error reading message: ", err)
            return
        }
            
        c.Pool.Broadcast <- message
    }
}

