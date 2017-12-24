package main

import (
	"math/rand"
	"time"
)

func init() {
	// Make sure that we can generate random numbers
	rand.Seed(time.Now().UnixNano())
}

func main() {
	startServer()
}
