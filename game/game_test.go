package game

import "testing"

func TestWelcomeMessage(t *testing.T) {
	if welcomeMessage != "Adventure game!" {
		t.Fatal("unexpected welcome message:", welcomeMessage)
	}
}

func TestNewGame(t *testing.T) {
	g := NewGame()

	if g.Player.Position != "kitchen" {
		t.Fatal("unexpected player starting position:", g.Player.Position)
	}
}

func TestGameComands(t *testing.T) {
	g := &Game{}

	if got, want := len(g.commands()), 10; got != want {
		t.Fatalf("unexpected number of game commands: %d, want %d", got, want)
	}
}

func TestGamePlace(t *testing.T) {
	g := NewGame()
	p := g.Place()

	if got, want := p.Name, "kitchen"; got != want {
		t.Fatal("unexpected starting place:", g.Player.Position)
	}

	if got, want := p.IsNextTo("garden"), true; got != want {
		t.Fatalf(`p.IsNextTo("garden") = %v, want %v`, got, want)
	}

	if got, want := p.IsNextTo("attic"), false; got != want {
		t.Fatalf(`p.IsNextTo("attic") = %v, want %v`, got, want)
	}
}

func TestGameGeneric(t *testing.T) {
	g := &Game{}

	out, err := g.generic("", []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != genericMessage {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestGameInventory(t *testing.T) {
	for i, tt := range []struct {
		items Items
		out   string
	}{
		{Items{}, notCarryingMessage},
		{Items{"foo": &Item{}, "bar": &Item{}}, "You are carrying: bar, foo"},
	} {
		g := &Game{Player: &Player{Inventory: tt.items}}

		out, err := g.inventory("", []string{})
		if err != nil {
			t.Fatalf("T%d: unexpected error: %v", i, err)
		}

		if out != tt.out {
			t.Fatalf("T%d: unexpected output: %q, want %q", i, out, tt.out)
		}
	}
}

func TestGameHelp(t *testing.T) {
	g := &Game{}

	out, err := g.help("", []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != youCanPerformMessage+
		"drop, exit, find, help, inventory, kill, look, take, use, walk" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestGameLook(t *testing.T) {
	for i, tt := range []struct {
		walks []string
		out   string
	}{
		{[]string{}, "You are standing in the kitchen for the first time.\nPaths: garden, livingroom"},
		{[]string{"garden"}, "You are standing in the garden for the first time.\nPaths: kitchen\nItems: ladder"},
	} {
		g := NewGame()

		for _, w := range tt.walks {
			g.walk("", []string{w})
		}

		out, err := g.look("", []string{})
		if err != nil {
			t.Fatalf("T%d: unexpected error: %v", i, err)
		}

		if out != tt.out {
			t.Fatalf("T%d: unexpected output: %q, want %q", i, out, tt.out)
		}
	}
}
