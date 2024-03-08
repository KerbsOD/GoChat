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
			break
		case message := <-pool.Broadcast:
			broadcastMessage(message, pool)
			break
		case client := <-pool.Unregister:
			unregisterClient(client, pool)
			break
		}
	}
}

func registerClient(client *Client, pool *Pool) {
	pool.Clients[client] = true
	for client := range pool.Clients {
		joinMessage := fmt.Sprintf("%s Connected...", client.ID)
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
	for client := range pool.Clients {
		leaveMessage := fmt.Sprintf("%s Disconnected...", client.ID)
		client.Conn.WriteJSON(Message{Type: 1, Sender: client.ID, Body: leaveMessage})
	}
}
