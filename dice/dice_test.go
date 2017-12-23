package dice

import "testing"

var d *Dice

func init(){
	d = NewDice()
}


func TestNewDice(t *testing.T) {

	if d.count != 2 {
		t.Fatalf("The count of dice should be 2\n")
	}

	if len(d.Thrown) != 0 {
		t.Fatalf("The Thrown property should be empty\n")
	}
}

func TestThrow(t *testing.T) {
	// Throw the dice
	d.Throw()

	if len(d.Thrown) != 2 {
		t.Fatalf("The Thrown property should be equal to length 2, got %d => %v\n", len(d.Thrown), d.Thrown)
	}

	for _, thrownNumber := range d.Thrown {
		if thrownNumber < 1 || thrownNumber > 6 {
			t.Fatalf("The thrown number is out of range %d\n", thrownNumber)
		}
	}
}