package game

import (
	"github.com/Dmitrev/gangsta-monopoly/board"
	"github.com/Dmitrev/gangsta-monopoly/dice"
	"github.com/Dmitrev/gangsta-monopoly/player"
)

type Game struct {
	Board    board.Board
	Dice     dice.Dice
	Players  []player.Player
	initDone bool
}

func NewGame() *Game {
	g := Game{}
	// Unit testing to see if the value has been initialized
	g.initDone = true
	return &g
}
