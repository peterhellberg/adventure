package game

import "testing"

func TestPathsString(t *testing.T) {
	for _, tt := range []struct {
		p Paths
		s string
	}{
		{Paths{"foo", "bar", "baz"}, "foo, bar, baz"},
		{Paths{"sit", "amet"}, "sit, amet"},
	} {
		if got := tt.p.String(); got != tt.s {
			t.Fatalf(`unexpected string: %q, want: %q`, got, tt.s)
		}
	}
}
