package game

import (
	"sort"
	"strings"
)

type Items map[string]*Item

func (i Items) String() string {
	s := []string{}

	for k, _ := range i {
		s = append(s, k)
	}

	sort.Strings(s)

	return strings.Join(s, ", ")
}

func (i Items) Set(name string, item *Item) {
	if i == nil {
		return
	}

	i[name] = item
}

type Item struct {
	Use    func(*Game) string
	Take   func(*Game) string
	Weight int
}

func newLadder() *Item {
	return &Item{
		Use: func(g *Game) string {
			p := g.Place()

			if p.Name == "livingroom" {
				if !contains(p.Paths, "attic") {
					p.Paths = append(p.Paths, "attic")
				}

				delete(g.Player.Items, "ladder")

				g.Player.Position = "attic"

				return "You used the ladder in order to reach the attic"
			}

			return "You can’t use the ladder here."
		},
		Take: func(g *Game) string {
			return "You took the ladder, it might be able to reach the attic"
		},
	}
}

func newCarrot() *Item {
	return &Item{Use: func(g *Game) string {
		delete(g.Player.Items, "carrot")
		return "You ate the carrot."
	}}
}

func newSpade() *Item {
	return &Item{}
}
