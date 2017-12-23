package player

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()

	if p.IsTurn || p.Name != "" {
		t.Fatalf("New player should return empty instance")
	}
}
