package game

type World struct {
	Places Places
}

func NewWorld() *World {
	return &World{Places: NewPlaces()}
}
