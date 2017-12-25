package main

import (
	"log"

	"github.com/Dmitrev/gangsta-monopoly/player"
	"github.com/gorilla/websocket"
)

// Handle incoming events from clients
func handleEvent(conn *websocket.Conn, msg Message) {
	// Before game start
	switch msg.Action {
	case "register":
		register(conn, msg.Data) // msg.Data will contain the name of the user
	}

	//if !g.Started() {
	//	return
	//}

	// During game
	switch msg.Action {
	case "throw_dice":
		throwDice(conn)
	}
}

// Send register request to client
// Asking to fill in name
func sendRegisterRequest(conn *websocket.Conn) {
	err := conn.WriteJSON(&struct {
		Action string `json:"action"`
	}{"register"})

	if err != nil {
		log.Printf("[register] failed to send register request (%s)\n", err)
	}
}

// Create a player and add it to the game
func register(conn *websocket.Conn, name string) {
	p := player.NewPlayer()
	p.Name = name
	p.Conn = conn

	log.Printf("Registering new user: %s", name)
	g.AddPlayer(p)

	conn.WriteJSON(&Message{"register_ok", ""})
}

func throwDice(conn *websocket.Conn) {
	p := g.GetPlayer(conn)
	log.Printf("%s throws the dice\n", p.Name)
	g.ThrowDice(p)
}
