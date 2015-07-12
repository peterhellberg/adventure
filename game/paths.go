package game

import "strings"

type Paths []string

func (p Paths) String() string {
	return strings.Join(p, ", ")
}
