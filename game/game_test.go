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

	g.walk("garden", "shed")
	g.take("ladder")
	g.walk("garden", "kitchen", "livingroom")
	g.use("ladder")

	if g.Player.Position != "attic" {
		t.Fatal("unexpected player position:", g.Player.Position)
	}

	g.walk("livingroom", "hallway", "basement")

	if g.Player.Position != "hallway" {
		t.Fatal("unexpected player position:", g.Player.Position)
	}

	g.walk("livingroom", "attic")
	g.take("flashlight")
	g.walk("livingroom", "hallway", "basement")

	if g.Player.Position != "basement" {
		t.Fatal("unexpected player position:", g.Player.Position)
	}

	out, err := g.look()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want := "YOU CAN SEE EVERYTHING\n\n" +
		"You are standing in the basement for the first time.\n" +
		"Paths: hallway"; out != want {
		t.Fatalf("unexpected output: %q, want %q", out, want)
	}

	g.drop("carrot")

	out, err = g.items()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if want := "You are carrying: flashlight"; out != want {
		t.Fatalf("unexpected output: %q, want %q", out, want)
	}
}

func TestGameComands(t *testing.T) {
	g := &Game{}

	if got, want := len(g.commands()), 8; got != want {
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

	out, err := g.generic()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != genericMessage {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestGameItems(t *testing.T) {
	for i, tt := range []struct {
		items Items
		out   string
	}{
		{Items{}, notCarryingMessage},
		{Items{"foo": &Item{}, "bar": &Item{}}, "You are carrying: bar, foo"},
	} {
		g := &Game{Player: &Player{Items: tt.items}}

		out, err := g.items()
		if err != nil {
			t.Fatalf("T%d: unexpected error: %v", i, err)
		}

		if out != tt.out {
			t.Fatalf("T%d: unexpected g.items output: %q, want %q", i, out, tt.out)
		}
	}
}

func TestGameHelp(t *testing.T) {
	g := &Game{}

	out, err := g.help()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != canPerformMessage+
		"drop, exit, help, items, look, take, use, walk" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestGameLook(t *testing.T) {
	for i, tt := range []struct {
		walks []string
		out   string
	}{
		{[]string{}, "You are standing in the kitchen for the first time.\nPaths: garden, livingroom"},
		{[]string{"garden", "shed"}, "You are standing in the shed for the first time.\nPaths: garden\nItems: ladder"},
	} {
		g := NewGame()

		g.walk(tt.walks...)

		out, err := g.look()
		if err != nil {
			t.Fatalf("T%d: unexpected error: %v", i, err)
		}

		if out != tt.out {
			t.Fatalf("T%d: unexpected output: %q, want %q", i, out, tt.out)
		}
	}
}

func TestGameExit(t *testing.T) {
	g := NewGame()

	out, err := g.exit()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != byeMessage {
		t.Fatalf("unexpected output: %q", out)
	}
}
