package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type          int    `json:"type"`
	StatusMessage int    `json:"statusmessage"`
	Sender        string `json:"sender"`
	Body          string `json:"body"`
}

type FrontMessage struct {
    Username string `json:"username"`
    Content string `json:"content"`
}

func (c *Client) Register(pool *Pool) {
	pool.Register <- c
}

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
            
        dataRecieved := 

		message := Message{Type: messageType, StatusMessage: 1, Sender: string(c.ID), Body: string(p)}
		c.Pool.Broadcast <- message
	}
}
