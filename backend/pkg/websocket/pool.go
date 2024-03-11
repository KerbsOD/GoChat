package websocket

import (
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			registerClient(client, pool)
		case message := <-pool.Broadcast:
			broadcastMessage(message, pool)
		case client := <-pool.Unregister:
			unregisterClient(client, pool)
		}
	}
}

func registerClient(client *Client, pool *Pool) {
	pool.Clients[client] = true

	log.Printf("User Connected: %+v\n", client.ID)

	for client := range pool.Clients {
		err := client.Conn.WriteJSON(Message{Type: 1, Sender: client.ID, Body: ""})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func broadcastMessage(message Message, pool *Pool) {
	log.Printf("Message Received: %+v\n", message)

	for client := range pool.Clients {
		err := client.Conn.WriteJSON(message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func unregisterClient(client *Client, pool *Pool) {
	delete(pool.Clients, client)

	log.Printf("User Disconnected: %+v\n", client.ID)

	for client := range pool.Clients {
		err := client.Conn.WriteJSON(Message{Type: 1, Sender: client.ID, Body: ""})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
