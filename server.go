package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Message struct {
	Action string
	Data   string
}

var clients = make(map[*websocket.Conn]bool)

func startServer() {
	// Create simple file serverinfo info wa to serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Setup handler for the websocket
	http.HandleFunc("/ws", websocketHandler)

	log.Printf("StartingServer server on port 8000\n")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("[websocketHandler] Failed to upgrade connection (%s)\n", err)
	}

	defer conn.Close()

	// Once the user gets here, that means we have a new connection
	// Add client to the clients slice
	clients[conn] = true
	log.Printf("New user connected\n")

	var msg Message
	// Connection loop
	for {
		// Incoming JSON from the client
		err := conn.ReadJSON(&msg)

		// If we cannot read the message anymore that means that the user is disconnected
		if err != nil {
			// Remove connection from clients
			delete(clients, conn)
			break
		}

	}

}
