package game

import (
	"errors"

	"github.com/Dmitrev/gangsta-monopoly/board"
	"github.com/Dmitrev/gangsta-monopoly/dice"
	"github.com/Dmitrev/gangsta-monopoly/player"
)

type Game struct {
	Board    board.Board
	Dice     dice.Dice
	Players  []*player.Player
	initDone bool
	started  bool
}

var ErrNotEnoughPlayers = errors.New("not enough players in game")
var ErrGameNotStarted = errors.New("game not started")

func NewGame() *Game {
	g := Game{}
	// Unit testing to see if the value has been initialized
	g.initDone = true
	return &g
}

func (g *Game) AddPlayer(p *player.Player) {
	g.Players = append(g.Players, p)

}

func (g *Game) StartGame() (err error) {
	if len(g.Players) < 2 {
		err = ErrNotEnoughPlayers
		return
	}
	g.started = true
	return
}

func (g *Game) FirstTurn() (err error) {
	if !g.started {
		err = ErrGameNotStarted
		return
	}

	g.Players[0].IsTurn = true
	return
}
