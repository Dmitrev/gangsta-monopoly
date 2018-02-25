package player

import "github.com/gorilla/websocket"

type Player struct {
	Id         int
	Name       string
	IsTurn     bool
	ThrownDice bool
	Conn       *websocket.Conn
	Position   int
	Ready      bool
	Money      int
	Send *chan []byte
}

func NewPlayer() *Player {
	return &Player{}
}

// Indicates if the player performed all the mandatory actions to end the turn
func (p *Player) AllowEndTurn() bool {
	return p.IsTurn && p.ThrownDice
}

// Perform clean up when ending turn
func (p *Player) EndTurn() {
	p.ThrownDice = false
}
