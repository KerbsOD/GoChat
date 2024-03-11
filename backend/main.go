package main

import (
	"fmt"
	"net/http"

	"github.com/KerbsOD/GoChat/pkg/websocket"
)

func main() {
	fmt.Println("===== GoChat v1.1 =====")
	pool := initializePool()
	setupRoutes(pool)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Server closed\n")
	}
}

func initializePool() *websocket.Pool {
	pool := websocket.NewPool()
	go pool.Start()
	return pool
}

// Cada vez que alguien accede al server mediante "8080/ws" se llama a serveWebSocket
func setupRoutes(pool *websocket.Pool) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(pool, w, r)
	})
}

func serveWebSocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		ID:   "Norber",
		Conn: conn,
		Pool: pool,
	}

	client.Register(pool)
	client.Read()
}
