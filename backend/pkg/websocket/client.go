package websocket

import (
	"encoding/json"
	"fmt"
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
	Content  string `json:"content"`
}

func (c *Client) Register(pool *Pool, conn *websocket.Conn) {
	log.Println("Sending username request")
	err := c.Conn.WriteJSON(Message{Type: 1, StatusMessage: 0, Sender: c.ID, Body: ""})
	if err != nil {
		log.Println(err)
		return
	}

	_, p, err := c.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(p))

	var messageWithName FrontMessage
	if err := json.Unmarshal(p, &messageWithName); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	username := messageWithName.Username

	c.ID = username

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

		message := Message{Type: messageType, StatusMessage: 2, Sender: c.ID, Body: string(p)}
		c.Pool.Broadcast <- message
	}
}
