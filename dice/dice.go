package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	count int
	Thrown []int
}

func NewDice() *Dice {
	d := Dice{2, []int{}}
	return &d
}

// Throw will generate new numbers in the Thrown slice
func (d *Dice) Throw() {
	// Make sure that we can generate random numbers
	rand.Seed(time.Now().UnixNano())
	max := 6
	min := 1
	thrown := make([]int, d.count)

	// Generate new numbers from 1-6
	for i := 0; i < d.count; i++ {
		random := rand.Intn(max - min) + min
		thrown = append(thrown, random)
	}

	// Set the Thrown property
	d.Thrown = thrown
}