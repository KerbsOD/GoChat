package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Each client has it's own websocket connection and a reference to the shared pool.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type   int    `json:"type"`
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

// When the websocket connection receives an array of bytes it broadcasts it to the rest of clients.
// c.Conn.ReadMessage() it's locked until it recieves data.
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Sender: string(c.ID), Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
