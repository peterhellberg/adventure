package game

import (
	"strings"

	"github.com/abiosoft/ishell"
)

type Game struct {
	*ishell.Shell
	Player *Player
	World  *World
}

const (
	WelcomeMessage = "Adventure game!"
)

func Start() {
	g := &Game{ishell.NewShell(), NewPlayer(), NewWorld()}

	g.Println(WelcomeMessage)
	g.Setup()
	g.Start()
}

func (g *Game) commands() map[string]func(string, []string) (string, error) {
	return map[string]func(string, []string) (string, error){
		"help": g.help,
		"look": g.look,
		"walk": g.walk,
		"take": g.take,
		"drop": g.drop,
		"exit": g.exit,
	}
}

func (g *Game) Setup() {
	g.SetPrompt("▶ ")
	g.RegisterGeneric(g.generic)

	for n, c := range g.commands() {
		g.Register(n, c)
	}
}

func (g *Game) Place() *Place {
	if p := g.World.Places[g.Player.Position]; p != nil {
		return p
	}

	return &Place{Name: "void", Paths: []string{"nowhere"}}
}

func (g *Game) generic(cmd string, args []string) (string, error) {
	return "I don’t know how to do that.", nil
}

func (g *Game) help(cmd string, args []string) (string, error) {
	commandNames := []string{}

	for n, _ := range g.commands() {
		commandNames = append(commandNames, n)
	}

	return "You can perform the following commands:\n" +
		strings.Join(commandNames, ", "), nil
}

func (g *Game) look(cmd string, args []string) (string, error) {
	p := g.Place()

	g.Println("You are standing in the " + p.Name)

	if len(p.Paths) > 0 {
		g.Println("Paths: " + p.Paths.String())
	}

	if len(p.Items) > 0 {
		g.Println("Items: " + p.Items.String())
	}

	return "", nil
}

func (g *Game) take(cmd string, args []string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to take.", nil
	}

	p := g.Place()

	if _, ok := p.Items[args[0]]; ok {
		return "You took it!", nil
	}

	return "You can’t take that, since it doesn’t exist.", nil
}

func (g *Game) drop(cmd string, args []string) (string, error) {
	return "You don’t know how to do that.", nil
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

		return "You walked to the " + target, nil
	}

	return "You can’t go there!", nil
}

func (g *Game) exit(cmd string, args []string) (string, error) {
	g.Stop()

	return "Good bye! Player 1", nil
}
