package dice

import (
	"math/rand"
)

type Dice struct {
	Count  int
	Thrown []int
}

// Throw will generate new numbers in the Thrown slice
func (d *Dice) Throw() {
	max := 6
	min := 1
	thrown := make([]int, d.Count)

	// Generate new numbers from 1-6
	for i := 0; i < d.Count; i++ {
		random := rand.Intn(max-min) + min
		thrown[i] = random
	}

	// Set the Thrown property
	d.Thrown = thrown
}

// Sum returns the sum of all the numbers from the dice
// Will be mainly used to determine the amount of steps you can take
func (d *Dice) Sum() int {
	total := 0
	for _, number := range d.Thrown {
		total += number
	}

	return total
}

// Double returns true if all the numbers thrown are the same
func (d *Dice) Double() bool {

	firstNumber := 0

	for _, number := range d.Thrown {
		// If first number is not set this is the first loop
		// And we will set the first number to number from the first loop
		if firstNumber == 0 {
			firstNumber = number
			continue
		}

		// If we get a mismatch the could never be a double throw so return
		if firstNumber != number {
			return false
		}
	}

	// if you get through the loop then it's a double
	return true
}
