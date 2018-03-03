package main

import (
	"log"

	"github.com/Dmitrev/gangsta-monopoly/player"
	"github.com/Dmitrev/gangsta-monopoly/networking"
)

type Register struct{
	Name string `json:"name"`
}

// Handle incoming events from clients
func handleEvent(client *Client, stream []byte) {
	msg, err := networking.DeserializeToMessage(stream)

	if err != nil {
		log.Printf("Problem deserializing incoming message %#v", stream)
	}

	// Before game start
	switch msg.Type {
	case "register":
		register(client, msg) // msg.Data will contain the name of the user
	case "ready":
		ready(client, msg)
	}

	//if !g.Started() {
	//	return
	//}

	// During game
	switch msg.Type {
	case "throw_dice":
		throwDice(client)
	}
}

// Send register request to client
// Asking to fill in name
func sendRegisterRequest(c *Client) {
	// Create new player
	p := player.NewPlayer()
	// Add connection to player to have a reference to client
	p.Conn = c.conn

	p.Send = &c.send

	// Add player to client
	c.player = p
	// Add player to game
	g.AddPlayer(p)

	// Send request to register
	// Create new message
	msg, err := networking.NewMessage("register",nil)
	if err != nil {
		return
	}

	// Serialize and send
	serialized, err := msg.Serialize()
	if err != nil {
		return
	}
	c.send <- serialized
}

// Create a player and add it to the game
func register(client *Client, msg *networking.Message) {

	var r Register
	log.Printf("%s", msg.Data)

	err := networking.DeserializeMessageData(msg, &r)

	if err != nil {
		log.Printf("Error deserializing message data %s", err)
	}

	p := client.player
	p.Name = r.Name

	log.Printf("Registering new user: %s", p.Name)

	// Send player information to all clients
	g.SendAllPlayersPositions()
	g.SendPlayerUpdate(p, "register_ok", nil)
}

func ready(client *Client, msg *networking.Message) {
	p := client.player
	// Toggle ready state
	p.Ready = !p.Ready
	g.CheckStartGame()

	// Send player information to all clients
	g.SendAllPlayersPositions()
}

func throwDice(client *Client) {
	p := client.player

	if !p.IsTurn || p.ThrownDice {
		log.Printf("%s tried to throw before his turn!", p.Name)
		return
	}

	log.Printf("%s throws the dice\n", p.Name)
	g.ThrowDice(p)
}
