package stringutil

import (
	"fmt"
	"testing"
)

func BenchmarkFindClosestMatches(b *testing.B) {
	candidates := []string{
		"copilot", "claude", "codex", "custom", "contents", "content", "context", "controls",
		"push", "pull_request", "issues", "actions", "checks", "secrets", "variables",
		"environment", "deploy", "build", "test", "lint", "release", "tag", "branch",
	}
	target := "copiliot"
	maxResults := 3

	b.ResetTimer()
	for range b.N {
		FindClosestMatches(target, candidates, maxResults)
	}
}

func BenchmarkLevenshteinDistance_Small(b *testing.B) {
	s1 := "copilot"
	s2 := "copiliot"

	b.ResetTimer()
	for range b.N {
		LevenshteinDistance(s1, s2)
	}
}

func BenchmarkLevenshteinDistance_Large(b *testing.B) {
	s1 := "extremely-long-string-to-test-performance-of-levenshtein-distance-algorithm"
	s2 := "extremely-long-string-to-test-performance-of-levenshtein-distance-algorith"

	b.ResetTimer()
	for range b.N {
		LevenshteinDistance(s1, s2)
	}
}

func BenchmarkFindClosestMatches_Large(b *testing.B) {
	candidates := make([]string, 1000)
	for i := range 1000 {
		candidates[i] = fmt.Sprintf("candidate-%d", i)
	}
	target := "candidate-999-typo"
	maxResults := 5

	b.ResetTimer()
	for range b.N {
		FindClosestMatches(target, candidates, maxResults)
	}
}
