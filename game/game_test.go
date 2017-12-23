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
