package game

import "fmt"

type Player struct {
	Name     string
	Position string
	Items    Items
}

func NewPlayer() *Player {
	return &Player{
		Position: "kitchen",
		Items: Items{
			"carrot": newCarrot(),
		},
	}
}
func (p *Player) Item(name string) (*Item, error) {
	if item, ok := p.Items[name]; ok {
		return item, nil
	}

	return &Item{}, fmt.Errorf("missing item")
}

func (p *Player) HasItem(name string) bool {
	if _, ok := p.Items[name]; ok {
		return true
	}

	return false
}

func (p *Player) AddItem(name string, item *Item) {
	if p.Items == nil {
		p.Items = Items{}
	}

	p.Items[name] = item
}
