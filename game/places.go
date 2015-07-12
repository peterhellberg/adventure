package game

type Places map[string]*Place

type Place struct {
	Name  string
	Paths Paths
	Items Items
}

var (
	kitchen = &Place{
		"kitchen",
		Paths{"garden", "livingroom"},
		Items{"foobar": &Item{}},
	}

	garden = &Place{
		"garden",
		Paths{"kitchen"},
		Items{},
	}

	livingroom = &Place{
		"livingroom",
		Paths{"kitchen", "attic"},
		Items{},
	}

	attic = &Place{
		"attic",
		Paths{"livingroom"},
		Items{"flashlight": &Item{}},
	}
)

func NewPlaces() Places {
	return Places{
		"kitchen":    kitchen,
		"garden":     garden,
		"livingroom": livingroom,
		"attic":      attic,
	}
}

func (p *Place) IsNextTo(name string) bool {
	return contains(p.Paths, name)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
