package game

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	if !g.initDone {
		t.Fatalf("Game did not init")
	}

}
