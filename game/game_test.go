package game

import (
	"testing"

	"github.com/Dmitrev/gangsta-monopoly/player"
)

func TestNewGame(t *testing.T) {
	g := NewGame()
	if !g.initDone {
		t.Fatalf("Game did not init")
	}

}

func TestGame_AddPlayer(t *testing.T) {
	g := NewGame()
	p := player.NewPlayer()
	g.AddPlayer(p)

	if len(g.Players) != 1 {
		t.Fatalf("Player count should be 1")
	}
}

func TestGame_StartGame(t *testing.T) {
	g := NewGame()
	p1 := player.NewPlayer()
	p2 := player.NewPlayer()
	g.AddPlayer(p1)
	g.AddPlayer(p2)

	err := g.StartGame()

	if err != nil {
		t.Fatalf("Game is not able to start, err: %s", err)
	}
}

func TestGame_StartGameReturnsNotEnoughPlayersError(t *testing.T) {
	g := NewGame()

	// Testing starting game with no players
	err := g.StartGame()

	if err != ErrNotEnoughPlayers {
		t.Fatalf("Game should return ErrNotEnoughPlayers error")
	}

	// Testing starting game with not enough players
	p := player.NewPlayer()
	g.AddPlayer(p)

	err = g.StartGame()

	if err != ErrNotEnoughPlayers {
		t.Fatalf("Game should return ErrNotEnoughPlayers error")
	}

}

func TestGame_FirstTurn(t *testing.T) {
	g := NewGame()

	p1 := player.NewPlayer()
	p2 := player.NewPlayer()
	g.AddPlayer(p1)
	g.AddPlayer(p2)

	err := g.StartGame() // StartGame() calls FirstTurn()

	if err != nil {
		t.Fatalf("Could not start game, err: %s", err)
	}

	err = g.FirstTurn()

	if err != nil {
		t.Fatalf("Could not set first turn, err: %s", err)
	}

	if !g.Players[0].IsTurn {
		t.Fatalf("FirstTurn did not properly set the IsTurn to true on the first player")
	}
}
