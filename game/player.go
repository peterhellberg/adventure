package game

type Player struct {
	Name      string
	Position  string
	Inventory Items
}

func NewPlayer() *Player {
	return &Player{
		Position: "kitchen",
		Inventory: Items{
			"carrot": newCarrot(),
		},
	}
}
