package game

import (
	"errors"

	"log"

	"fmt"

	"github.com/Dmitrev/gangsta-monopoly/board"
	"github.com/Dmitrev/gangsta-monopoly/dice"
	"github.com/Dmitrev/gangsta-monopoly/networking"
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

type Update struct {
	Receiver *player.Player
	Message  *networking.Message
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
	g.SendAllPlayersUpdate("game_started", nil)
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
	if g.currentTurn > -1 {
		currentPlayer := g.Players[g.currentTurn]
		currentPlayer.EndTurn()
	}

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
	g.SendAllPlayersUpdate("next_turn", struct {
		Name string `json:"name"`
	}{nextPlayer.Name})

	g.SendPlayerUpdate(nextPlayer, "your_turn", nil)
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
	// Generate dice numbers
	g.Dice.Throw()
	// Indicate that the player has thrown
	p.ThrownDice = true
	g.SendPlayerUpdate(p, "thrown_dice", nil)
	log.Printf("%s threw (%v) %d", p.Name, g.Dice.Thrown, g.Dice.Sum())
	// Calculate next position
	nextPosition := g.NextPosition(p)
	log.Printf("%s's next position is %d", p.Name, nextPosition)
	// Move player to next position
	p.Position = nextPosition
	// Notify all players about the position change
	g.SendAllPlayersPositions()
}

func (g *Game) SendAllPlayersPositions() {

	var positions = make([]PositionUpdate, 0)

	// Collect all the positions
	for index, p := range g.Players {
		hasTurn := index == g.currentTurn
		positions = append(positions, PositionUpdate{p.Name, p.Position, p.Ready, hasTurn})
	}
	// Send all the positions
	g.SendAllPlayersUpdate("position_update", positions)
}

func (g *Game) BroadCastPlayerLeft(p *player.Player) {
	// Send all the positions
	message := fmt.Sprintf("%s left the game", p.Name)
	g.SendAllPlayersUpdate("player_left", message)
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

func (g *Game) SendPlayerUpdate(p *player.Player, messageType string, data interface{}) {
	msg, err := networking.NewMessage(messageType, data)

	if err != nil {
		log.Fatal(err)
	}

	serialized, err := msg.Serialize()

	if err != nil {
		return
	}

	*p.Send <- serialized
}
func (g *Game) SendAllPlayersUpdate(messageType string, data interface{}) {
	msg, err := networking.NewMessage(messageType, data)

	if err != nil {
		log.Fatal(err)
	}

	serialized, err := msg.Serialize()

	if err != nil {
		return
	}

	for _, p := range g.Players {
		log.Printf("Send update to player %s", p.Name)
		*p.Send <- serialized
	}
}
