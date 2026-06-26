package stringutil

import (
	"fmt"
	"testing"
)

func BenchmarkLevenshteinDistance(b *testing.B) {
	cases := []struct {
		s1, s2 string
	}{
		{"", ""},
		{"a", "b"},
		{"kitten", "sitting"},
		{"flaw", "lawn"},
		{"distance", "difference"},
		{"a very long string that might be compared", "another very long string for comparison"},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%d-%d", len(tc.s1), len(tc.s2))
		b.Run(name, func(b *testing.B) {
			for range b.N {
				LevenshteinDistance(tc.s1, tc.s2)
			}
		})
	}
}
