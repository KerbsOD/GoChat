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

func (p *Pool) Run() {
	for {
		select {
		case client := <-p.Register:
			registerClient(client, p)
		case message := <-p.Broadcast:
			broadcastMessage(message, p)
		case client := <-p.Unregister:
			unregisterClient(client, p)
		}
	}
}

func registerClient(client *Client, pool *Pool) {
	pool.Clients[client] = true
	message := Message{Type: 1, Sender: client.ID, Body: ""}

	log.Println("User Connected: ", client.ID)

	for client := range pool.Clients {
		if err := client.SendMessage(message); err != nil {
			log.Println(err)
			return
		}
	}
}

func broadcastMessage(message Message, pool *Pool) {
	for client := range pool.Clients {
		log.Printf("Message {Type: %v, Sender: %v, Content: %v} ", message.Type, message.Sender, message.Body)
		if err := client.SendMessage(message); err != nil {
			log.Println(err)
			return
		}
	}
}

func unregisterClient(client *Client, pool *Pool) {
	delete(pool.Clients, client)
	message := Message{Type: 3, Sender: client.ID, Body: ""}

	log.Println("User Disconnected: ", client.ID)

	for client := range pool.Clients {
		if err := client.SendMessage(message); err != nil {
			log.Println(err)
			return
		}
	}
}
