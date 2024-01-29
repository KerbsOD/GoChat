package main

import (
	"fmt"
	"net/http"

	"github.com/KerbsOD/GoChat/pkg/websocket"
)

// We initialize our channels via the NewPool() function
// The pool handling must be concurrent with the TCP listener in main.
func initializePool() *websocket.Pool {
	pool := websocket.NewPool()
	go pool.Start()
	return pool
}

// Registration of the patterns with their corresponding function handler.
func setupRoutes(pool *websocket.Pool) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

// We upgrade the CLient's HTTP connection to Websocket protocol and register it onto the pool.
// The client now will have access to broadcasted messages until it's disconnected.
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		ID:   "Octo",
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func main() {
	fmt.Println("===== GoChat v1.0 =====")
	pool := initializePool()
	setupRoutes(pool)
	http.ListenAndServe(":8080", nil)
}
