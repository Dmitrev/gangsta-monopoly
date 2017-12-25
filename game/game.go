package game

import (
	"errors"

	"log"

	"github.com/Dmitrev/gangsta-monopoly/board"
	"github.com/Dmitrev/gangsta-monopoly/dice"
	"github.com/Dmitrev/gangsta-monopoly/player"
	"github.com/gorilla/websocket"
)

type Game struct {
	Board    board.Board
	Dice     *dice.Dice
	Players  []*player.Player
	initDone bool
	started  bool
}

type PositionUpdate struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
	Ready    bool   `json:"ready"`
}

var ErrNotEnoughPlayers = errors.New("not enough players in game")
var ErrGameNotStarted = errors.New("game not started")

func NewGame() *Game {
	g := Game{}
	g.Dice = dice.NewDice()
	g.Board = board.NewBoard()
	// Unit testing to see if the value has been initialized
	g.initDone = true
	return &g
}

func (g *Game) Started() bool {
	return g.started
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

func (g *Game) GetPlayer(conn *websocket.Conn) *player.Player {
	for _, p := range g.Players {
		if conn == p.Conn {
			return p
		}
	}

	return nil
}

func (g *Game) NextPosition(p *player.Player) int {
	newPosition := p.Position + g.Dice.Sum()

	log.Printf("[NextPosition] length board: %d", len(g.Board.Spaces))
	log.Printf("[NextPosition] calculated new position: %d", newPosition)

	// Check if the new position is withing range of the board
	if newPosition <= len(g.Board.Spaces) {
		return newPosition
	}

	// If not, calculate how much we are beyond the index
	overflow := newPosition - len(g.Board.Spaces)
	log.Printf("[NextPosition] overflow: %d", newPosition)

	// Because indexes start at 0 we need to compensate for that
	return overflow - 1

}

func (g *Game) ThrowDice(p *player.Player) {

	log.Printf("[g.ThrowDice] %#v", g.Dice)
	g.Dice.Throw()
	log.Printf("%s threw %d", p.Name, g.Dice.Sum())
	nextPosition := g.NextPosition(p)
	log.Printf("%s's next position is %d", p.Name, nextPosition)
	p.Position = nextPosition
	g.SendAllPlayersPositions()

}

func (g *Game) SendAllPlayersPositions() {

	var positions = make([]PositionUpdate, 0)

	// Collect all the positions
	for _, p := range g.Players {
		positions = append(positions, PositionUpdate{p.Name, p.Position, p.Ready})
	}

	// Send all the positions
	for _, p := range g.Players {
		p.Conn.WriteJSON(&struct {
			Action string           `json:"action"`
			Data   []PositionUpdate `json:"data"`
		}{"position_update", positions})
	}
}
