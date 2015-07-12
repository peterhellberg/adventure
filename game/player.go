package game

type Player struct {
	Name      string
	Position  string
	Inventory []string
}

func NewPlayer() *Player {
	return &Player{
		Position:  "kitchen",
		Inventory: []string{},
	}
}
