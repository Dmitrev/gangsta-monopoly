package player

import "github.com/gorilla/websocket"

type Player struct {
	Name     string
	IsTurn   bool
	Conn     *websocket.Conn
	Position int
}

func NewPlayer() *Player {
	return &Player{}
}
