package board

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	if len(b.Spaces) != 40 {
		t.Fatalf("Board should have 40 spaces")
	}
}
