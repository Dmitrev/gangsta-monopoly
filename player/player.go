package player

import "github.com/gorilla/websocket"

type Player struct {
	Id       int
	Name     string
	IsTurn   bool
	Conn     *websocket.Conn
	Position int
	Ready    bool
	Money    int
}

func NewPlayer() *Player {
	return &Player{}
}
