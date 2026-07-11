// Package stringutil provides utility functions for working with strings.
package stringutil

import (
	"slices"
	"strings"

	"github.com/github/gh-aw/pkg/logger"
)

var fuzzyMatchLog = logger.New("stringutil:fuzzy_match")

type match struct {
	value    string
	distance int
}

// FindClosestMatches finds the closest matching strings using Levenshtein distance.
// It returns up to maxResults matches that have a Levenshtein distance of 3 or less.
// Results are sorted by distance (closest first), then alphabetically for ties.
func FindClosestMatches(target string, candidates []string, maxResults int) []string {
	fuzzyMatchLog.Printf("FindClosestMatches: target=%q, candidates=%d, maxResults=%d", target, len(candidates), maxResults)
	const maxDistance = 3
	var matches []match
	targetLower := strings.ToLower(target)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if targetLower == candidateLower {
			continue
		}
		if absDiff(len(targetLower), len(candidateLower)) > maxDistance {
			continue
		}
		if d := LevenshteinDistance(targetLower, candidateLower); d <= maxDistance {
			matches = append(matches, match{value: candidate, distance: d})
		}
	}

	sortMatches(matches)
	results := limitResults(matches, maxResults)
	fuzzyMatchLog.Printf("FindClosestMatches: returning %d match(es) within distance %d", len(results), maxDistance)
	return results
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sortMatches(matches []match) {
	slices.SortFunc(matches, func(a, b match) int {
		if a.distance != b.distance {
			return a.distance - b.distance
		}
		return strings.Compare(a.value, b.value)
	})
}

func limitResults(matches []match, maxResults int) []string {
	n := min(len(matches), maxResults)
	if n <= 0 {
		return nil
	}
	results := make([]string, n)
	for i := range n {
		results[i] = matches[i].value
	}
	return results
}

// LevenshteinDistance computes the Levenshtein distance between two strings.
// This is the minimum number of single-character edits (insertions, deletions, or substitutions)
// required to change one string into the other.
func LevenshteinDistance(a, b string) int {
	// Early exit for empty strings
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}

	// Ensure b is the shorter string to minimize DP table size
	if len(a) < len(b) {
		a, b = b, a
	}
	aLen, bLen := len(a), len(b)

	// Use a stack-allocated buffer for common small strings to avoid heap allocation.
	// 64 is a reasonable limit for typical identifiers or command names.
	var buffer [65]int
	var row []int
	if bLen+1 <= len(buffer) {
		row = buffer[:bLen+1]
	} else {
		row = make([]int, bLen+1)
	}

	// Initialize the first row (distance from empty string)
	for j := 0; j <= bLen; j++ {
		row[j] = j
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		prevSub := row[0] // previousRow[j-1]
		row[0] = i        // currentRow[0]

		for j := 1; j <= bLen; j++ {
			prevDel := row[j]   // previousRow[j]
			prevIns := row[j-1] // currentRow[j-1]

			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			row[j] = min(prevDel+1, min(prevIns+1, prevSub+cost))
			prevSub = prevDel
		}
	}

	return row[bLen]
}
