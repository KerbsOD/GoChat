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
	http.ListenAndServe(":8080", nil)
}

func initializePool() *websocket.Pool {
	pool := websocket.NewPool()
	go pool.Start()
	return pool
}

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

	pool.Register <- client
	client.Read()
}
