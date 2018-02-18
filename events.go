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
	case "ready":
		ready(conn)
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
	p := player.NewPlayer()
	p.Conn = conn
	g.AddPlayer(p)
	g.SendPlayerUpdate(p, "register", nil)
}

// Create a player and add it to the game
func register(conn *websocket.Conn, name string) {
	p, index := g.GetPlayer(conn)
	if index == -1 {
		log.Fatal("Trying to register undefined user")
	}
	p.Name = name

	log.Printf("Registering new user: %s", name)

	// Send player information to all clients
	g.SendAllPlayersPositions()
	g.SendPlayerUpdate(p, "register_ok", nil)
}

func ready(conn *websocket.Conn) {
	p, _ := g.GetPlayer(conn)
	// Toggle ready state
	p.Ready = !p.Ready
	g.CheckStartGame()

	// Send player information to all clients
	g.SendAllPlayersPositions()
}

func throwDice(conn *websocket.Conn) {
	p, _ := g.GetPlayer(conn)

	if !p.IsTurn || p.ThrownDice {
		log.Printf("%s tried to throw before his turn!", p.Name)
		return
	}

	log.Printf("%s throws the dice\n", p.Name)
	g.ThrowDice(p)
}
