package game

import (
	"errors"

	"log"

	"fmt"

	"github.com/Dmitrev/gangsta-monopoly/board"
	"github.com/Dmitrev/gangsta-monopoly/dice"
	"github.com/Dmitrev/gangsta-monopoly/player"
	"github.com/gorilla/websocket"
)

const (
	GameCreating = iota
	GamePlaying
	GameFinished
)

type Game struct {
	Status      int
	Board       *board.Board
	Dice        *dice.Dice
	Players     []*player.Player
	initDone    bool
	started     bool
	currentTurn int
}

type PositionUpdate struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
	Ready    bool   `json:"ready"`
	HasTurn  bool   `json:"has_turn"`
}

var ErrNotEnoughPlayers = errors.New("not enough players in game")
var ErrGameNotStarted = errors.New("game not started")

func NewGame() *Game {
	g := Game{
		Status: 0,
		Board: &board.Board{
			Spaces: make([]*board.Space, 40),
		},
		Dice:        &dice.Dice{Count: 2, Thrown: []int{}},
		Players:     nil,
		initDone:    false,
		started:     false,
		currentTurn: 0,
	}
	// Unit testing to see if the value has been initialized
	g.initDone = true
	return &g
}

//join
//leave
//playTurn
//start

func (g *Game) Started() bool {
	return g.started
}

func (g *Game) AddPlayer(p *player.Player) {
	g.Players = append(g.Players, p)

}

func (g *Game) CheckStartGame() {

	if len(g.Players) < 2 {
		log.Printf("Not enough players to start")
		return
	}

	// Check if all players are ready
	allReady := true

	for _, p := range g.Players {
		if !p.Ready {
			allReady = false
		}
	}

	if !allReady {
		log.Printf("Not all players are ready")
		return
	}

	// Game is ready to start!
	g.StartGame()
}

func (g *Game) StartGame() (err error) {
	if len(g.Players) < 2 {
		err = ErrNotEnoughPlayers
		return
	}
	g.started = true
	g.updateAll("game_started", nil)
	g.FirstTurn()
	return
}

func (g *Game) FirstTurn() (err error) {
	if !g.started {
		err = ErrGameNotStarted
		return
	}

	g.currentTurn = -1
	g.NextTurn()
	return
}

func (g *Game) NextTurn() {

	nextTurn := g.currentTurn + 1

	// Out of bounds
	if nextTurn > len(g.Players)-1 {
		nextTurn = 0
	}

	// Reset all player's turns
	for _, p := range g.Players {
		p.IsTurn = false
	}

	g.currentTurn = nextTurn
	nextPlayer := g.Players[nextTurn]
	nextPlayer.IsTurn = true
	log.Printf("%s is not on turn", nextPlayer.Name)
	g.updateAll("next_turn", struct {
		Name string `json:"name"`
	}{nextPlayer.Name})

	g.updatePlayer(nextPlayer, "your_turn", nil)
}

func (g *Game) GetPlayer(conn *websocket.Conn) (*player.Player, int) {
	for index, p := range g.Players {
		if conn == p.Conn {
			return p, index
		}
	}

	return nil, -1
}

func (g *Game) NextPosition(p *player.Player) int {
	newPosition := p.Position + g.Dice.Sum()

	log.Printf("[NextPosition] length board: %d", len(g.Board.Spaces))
	log.Printf("[NextPosition] calculated new position: %d", newPosition)

	// Check if the new position is withing range of the board
	if newPosition <= len(g.Board.Spaces)-1 {
		return newPosition
	}

	// If not, calculate how much we are beyond the index
	overflow := newPosition - (len(g.Board.Spaces) - 1)
	log.Printf("[NextPosition] calculate overflow %d - %d - 1 = %d", newPosition, len(g.Board.Spaces), overflow)
	log.Printf("[NextPosition] overflow: %d", overflow)

	// Because indexes start at 0 we need to compensate for that
	return overflow - 1

}

func (g *Game) ThrowDice(p *player.Player) {

	log.Printf("[g.ThrowDice] %#v", g.Dice)
	g.Dice.Throw()
	log.Printf("%s threw (%v) %d", p.Name, g.Dice.Thrown, g.Dice.Sum())
	nextPosition := g.NextPosition(p)
	log.Printf("%s's next position is %d", p.Name, nextPosition)
	p.Position = nextPosition
	g.SendAllPlayersPositions()
	g.NextTurn()

}

func (g *Game) SendAllPlayersPositions() {

	var positions = make([]PositionUpdate, 0)

	// Collect all the positions
	for index, p := range g.Players {
		hasTurn := index == g.currentTurn
		positions = append(positions, PositionUpdate{p.Name, p.Position, p.Ready, hasTurn})
	}

	// Send all the positions
	for _, p := range g.Players {
		p.Conn.WriteJSON(&struct {
			Action string           `json:"action"`
			Data   []PositionUpdate `json:"data"`
		}{"position_update", positions})
	}
}

func (g *Game) BroadCastPlayerLeft(p *player.Player) {
	// Send all the positions
	message := fmt.Sprintf("%s left the game", p.Name)

	for _, p := range g.Players {
		p.Conn.WriteJSON(&struct {
			Action string `json:"action"`
			Data   string `json:"data"`
		}{"player_left", message})
	}
}

func (g *Game) RemovePlayerByIndex(index int) {
	g.Players = append(g.Players[:index], g.Players[index+1:]...)
}

func (g *Game) Disconnect(conn *websocket.Conn) {
	p, index := g.GetPlayer(conn)
	// Remvoe player from the players slice
	g.RemovePlayerByIndex(index)
	// Let other players know who left
	g.BroadCastPlayerLeft(p)
	// Update the players list
	g.SendAllPlayersPositions()

}

func (g *Game) updateAll(action string, data interface{}) {
	for _, p := range g.Players {
		g.updatePlayer(p, action, data)
	}
}

func (g *Game) updatePlayer(p *player.Player, action string, data interface{}) {

	p.Conn.WriteJSON(&struct {
		Action string      `json:"action"`
		Data   interface{} `json:"data"`
	}{action, data})

}
