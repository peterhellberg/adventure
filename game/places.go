package game

import (
	"fmt"
	"strings"
)

type Places map[string]*Place

type Place struct {
	Name       string
	Paths      Paths
	Items      Items
	VisitCount int
	Enter      func(*Game) (string, error)
	Look       func(*Game) (string, error)
}

func NewPlaces() Places {
	return Places{
		"attic": &Place{
			Name:  "attic",
			Paths: Paths{"livingroom"},
			Items: Items{"flashlight": &Item{}},

			VisitCount: 1,
		},

		"basement": &Place{
			Name:  "basement",
			Paths: Paths{"hallway"},
			Items: Items{},

			Enter: func(g *Game) (string, error) {
				if g.Player.Position == "hallway" && !g.Player.HasItem("flashlight") {
					return "", fmt.Errorf("It is very dark in the basement, you are not be able to see anything.")
				}

				return "", nil
			},

			Look: func(g *Game) (string, error) {
				if g.Player.HasItem("flashlight") {
					return "YOU CAN SEE EVERYTHING", nil
				}

				return "", fmt.Errorf("It is very dark in the basement, you are not be able to see anything.")
			},
		},

		"bedroom": &Place{
			Name:  "bedroom",
			Paths: Paths{"hallway"},
			Items: Items{},
		},

		"forest": &Place{
			Name:  "forest",
			Paths: Paths{"garden"},
			Items: Items{},
		},

		"garden": &Place{
			Name:  "garden",
			Paths: Paths{"kitchen", "shed", "forest"},
			Items: Items{"spade": newSpade()},
		},

		"hallway": &Place{
			Name:  "hallway",
			Paths: Paths{"livingroom", "basement", "bedroom"},
			Items: Items{},
		},

		"kitchen": &Place{
			Name:  "kitchen",
			Paths: Paths{"garden", "livingroom"},
			Items: Items{},

			VisitCount: 1,
		},

		"livingroom": &Place{
			Name:  "livingroom",
			Paths: Paths{"kitchen", "hallway"},
			Items: Items{},
		},

		"shed": &Place{
			Name:  "shed",
			Paths: Paths{"garden"},
			Items: Items{"ladder": newLadder()},
		},
	}
}

func (p *Place) AddItem(name string, item *Item) {
	if p.Items == nil {
		p.Items = Items{}
	}

	p.Items[name] = item
}

func (p *Place) IsNextTo(name string) bool {
	return contains(p.Paths, name)
}

func (p *Place) describe() string {
	l := []string{}

	switch p.VisitCount {
	case 1:
		l = append(l, "You are standing in the "+p.Name+" for the first time.")
	default:
		l = append(l, "You are standing in the "+p.Name+fmt.Sprintf(" count: %d", p.VisitCount))
	}

	if len(p.Paths) > 0 {
		l = append(l, "Paths: "+p.Paths.String())
	}

	if len(p.Items) > 0 {
		l = append(l, "Items: "+p.Items.String())
	}

	return strings.Join(l, "\n")
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
