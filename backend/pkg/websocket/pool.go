package websocket

import "fmt"

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

	joinMessage := fmt.Sprintf("%s connected...", client.ID)
	fmt.Println(joinMessage)

	for client := range pool.Clients {
		client.Conn.WriteJSON(Message{Type: 1, Sender: client.ID, Body: joinMessage})
	}
}

func broadcastMessage(message Message, pool *Pool) {
	for client := range pool.Clients {
		err := client.Conn.WriteJSON(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func unregisterClient(client *Client, pool *Pool) {
	delete(pool.Clients, client)

	leaveMessage := fmt.Sprintf("%s disconnected...", client.ID)
	fmt.Println(leaveMessage)

	for client := range pool.Clients {
		client.Conn.WriteJSON(Message{Type: 1, Sender: client.ID, Body: leaveMessage})
	}
}
