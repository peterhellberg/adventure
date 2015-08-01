package game

import "fmt"

// Player contains the name of the player, the position and carried items
type Player struct {
	Name     string
	Position string
	Items    Items
}

// NewPlayer creates a new player
func NewPlayer() *Player {
	return &Player{
		Position: "kitchen",
		Items: Items{
			"carrot": newCarrot(),
		},
	}
}

// Item returns the item with the given name
func (p *Player) Item(name string) (*Item, error) {
	if item, ok := p.Items[name]; ok {
		return item, nil
	}

	return &Item{}, fmt.Errorf("missing item")
}

// HasItem returns true if the player holds the item
func (p *Player) HasItem(name string) bool {
	if _, ok := p.Items[name]; ok {
		return true
	}

	return false
}

// AddItem adds the item to the players inventory
func (p *Player) AddItem(name string, item *Item) {
	if p.Items == nil {
		p.Items = Items{}
	}

	p.Items[name] = item
}
