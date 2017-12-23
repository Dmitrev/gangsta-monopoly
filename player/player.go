package player

type Player struct{
	Name string
	IsTurn bool
}

func NewPlayer() *Player{
	return &Player{}
}