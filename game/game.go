package game

import (
	"sort"
	"strings"

	"github.com/abiosoft/ishell"
)

const (
	welcomeMessage       = "Adventure game!"
	genericMessage       = "You don’t know how to do that."
	notCarryingMessage   = "You are not carrying anything."
	carryingMessage      = "You are carrying: "
	youCanPerformMessage = "You can perform the following commands:\n"
)

type Game struct {
	*ishell.Shell
	Player *Player
	Places Places
}

func NewGame() *Game {
	g := &Game{ishell.NewShell(), NewPlayer(), NewPlaces()}

	g.Setup()

	return g
}

func Start() {
	g := NewGame()

	g.Println(welcomeMessage)
	g.Start()
}

func (g *Game) commands() map[string]func(string, []string) (string, error) {
	return map[string]func(string, []string) (string, error){
		"drop":      g.drop,
		"exit":      g.exit,
		"find":      g.generic,
		"help":      g.help,
		"inventory": g.inventory,
		"kill":      g.generic,
		"look":      g.look,
		"take":      g.take,
		"use":       g.use,
		"walk":      g.walk,
	}
}

func (g *Game) Setup() {
	g.SetPrompt("▶ ")
	g.RegisterGeneric(g.generic)

	for n, c := range g.commands() {
		g.Register(n, c)

		if n != "exit" {
			g.Register(string(n[0]), c)
		}
	}
}

func (g *Game) Place() *Place {
	if p := g.Places[g.Player.Position]; p != nil {
		return p
	}

	return &Place{Name: "void", Paths: []string{"nowhere"}}
}

func (g *Game) generic(cmd string, args []string) (string, error) {
	return genericMessage, nil
}

func (g *Game) inventory(cmd string, args []string) (string, error) {
	items := g.Player.Inventory

	if len(items) == 0 {
		return notCarryingMessage, nil
	}

	return carryingMessage + items.String(), nil
}

func (g *Game) help(cmd string, args []string) (string, error) {
	commandNames := []string{}

	for n, _ := range g.commands() {
		commandNames = append(commandNames, n)
	}

	sort.Strings(commandNames)

	return youCanPerformMessage + strings.Join(commandNames, ", "), nil
}

func (g *Game) look(cmd string, args []string) (string, error) {
	p := g.Place()

	l := []string{}

	switch p.VisitCount {
	case 1:
		l = append(l, "You are standing in the "+p.Name+" for the first time.")
	default:
		l = append(l, "You are standing in the "+p.Name)
	}

	if len(p.Paths) > 0 {
		l = append(l, "Paths: "+p.Paths.String())
	}

	if len(p.Items) > 0 {
		l = append(l, "Items: "+p.Items.String())
	}

	return strings.Join(l, "\n"), nil
}

func (g *Game) take(cmd string, args []string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to take.", nil
	}

	l := []string{}

	for _, name := range args {
		p := g.Place()

		if item, ok := p.Items[name]; ok {
			delete(p.Items, name)

			g.Player.Inventory[name] = item

			if item.Take != nil {
				return item.Take(g), nil
			}

			l = append(l, "You took the "+name)
		}
	}

	if len(l) > 0 {
		return strings.Join(l, "\n"), nil
	}

	return "You can’t take that which doesn’t exist.", nil
}

func (g *Game) drop(cmd string, args []string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to drop.", nil
	}

	l := []string{}

	for _, name := range args {
		p := g.Place()

		if item, ok := g.Player.Inventory[name]; ok {
			delete(g.Player.Inventory, name)

			p.Items[name] = item

			l = append(l, "You dropped the "+name)
		}
	}

	if len(l) > 0 {
		return strings.Join(l, "\n"), nil
	}

	return "Unable to drop something you are not carrying.", nil
}

func (g *Game) use(cmd string, args []string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to use.", nil
	}

	name := args[0]

	if item, ok := g.Player.Inventory[name]; ok {
		if item.Use != nil {
			return item.Use(g), nil
		}

		return "You can’t use the " + name, nil
	}

	return "You are not carrying that item.", nil
}

func (g *Game) walk(cmd string, args []string) (string, error) {
	if len(args) == 0 {
		return "You need to specify where to go.", nil
	}

	target := strings.ToLower(args[0])

	if len(args) > 2 && args[0] == "to" && args[1] == "the" {
		target = strings.ToLower(args[2])
	} else if len(args) > 1 && args[0] == "to" {
		target = strings.ToLower(args[1])
	}

	p := g.Place()

	if p.Name == target {
		return "You are already in the " + p.Name, nil
	}

	if p.IsNextTo(target) {
		g.Player.Position = target
		g.Places[target].VisitCount++

		return "You walked to the " + target, nil
	}

	return "You can’t go there!", nil
}

func (g *Game) exit(cmd string, args []string) (string, error) {
	g.Stop()

	return "Good bye! Player 1", nil
}
