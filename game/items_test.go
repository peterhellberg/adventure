package game

import "testing"

func TestItemsString(t *testing.T) {
	for _, tt := range []struct {
		i Items
		s string
	}{
		{Items{"foo": &Item{}, "bar": &Item{}, "baz": &Item{}}, "bar, baz, foo"},
		{Items{"sit": &Item{}, "amet": &Item{}}, "amet, sit"},
	} {
		if got := tt.i.String(); got != tt.s {
			t.Fatalf(`unexpected string: %q, want: %q`, got, tt.s)
		}
	}
}
