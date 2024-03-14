package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KerbsOD/GoChat/pkg/websocket"
)

func main() {
	fmt.Println(" ==== GoChat ==== ")

	pool := websocket.NewPool()
	go pool.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveNewClient(pool, w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Server closed: ", err)
	}
}

func serveNewClient(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println("Failed to upgrade connection: ", err)
	}

	client := websocket.NewClient(conn, pool)
	client.RequestUsername()
	pool.Register <- client
	client.Listen()
}
