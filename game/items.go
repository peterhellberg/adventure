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

type Item struct {
}
