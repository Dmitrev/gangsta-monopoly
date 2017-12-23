package board

type space struct {
	Name string
}

type Board struct {
	Spaces [40]space
}

func NewBoard() Board {
	var b Board
	b.Spaces = [40]space{}

	return b
}
