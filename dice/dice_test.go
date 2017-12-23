package dice

import "testing"

var d *Dice

func init(){
	d = NewDice()
}


func TestNewDice(t *testing.T) {

	if d.count != 2 {
		t.Fatalf("The count of dice should be 2")
	}

	if len(d.Thrown) != 0 {
		t.Fatalf("The Thrown property should be empty")
	}
}

func TestThrow(t *testing.T) {
	// Throw the dice
	d.Throw()

	if len(d.Thrown) != 2 {
		t.Fatalf("The Thrown property should be equal to length 2")
	}
}