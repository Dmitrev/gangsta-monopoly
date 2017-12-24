package main

import (
	"math/rand"
	"time"

	"github.com/Dmitrev/gangsta-monopoly/game"
)

func init() {
	// Make sure that we can generate random numbers
	rand.Seed(time.Now().UnixNano())
}

var g *game.Game

func main() {
	g = game.NewGame()
	startServer()
}
