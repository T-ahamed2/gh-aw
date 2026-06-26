// Package stringutil provides utility functions for working with strings.
package stringutil

import (
	"slices"
	"strings"

	"github.com/github/gh-aw/pkg/logger"
)

var fuzzyMatchLog = logger.New("stringutil:fuzzy_match")

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// FindClosestMatches finds the closest matching strings using Levenshtein distance.
// It returns up to maxResults matches that have a Levenshtein distance of 3 or less.
// Results are sorted by distance (closest first), then alphabetically for ties.
//
// This function is useful for "Did you mean?" suggestions when a user provides
// an unrecognized value (e.g., a typo in an engine name or event type).
func FindClosestMatches(target string, candidates []string, maxResults int) []string {
	fuzzyMatchLog.Printf("FindClosestMatches: target=%q, candidates=%d, maxResults=%d", target, len(candidates), maxResults)
	type match struct {
		value    string
		distance int
	}

	const maxDistance = 3 // Maximum acceptable Levenshtein distance

	var matches []match
	targetLower := strings.ToLower(target)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		// Skip exact matches
		if targetLower == candidateLower {
			continue
		}

		// Optimization: if the length difference is greater than maxDistance,
		// the Levenshtein distance must also be greater than maxDistance.
		if abs(len(targetLower)-len(candidateLower)) > maxDistance {
			continue
		}

		distance := LevenshteinDistance(targetLower, candidateLower)

		// Only include if distance is within acceptable range
		if distance <= maxDistance {
			matches = append(matches, match{value: candidate, distance: distance})
		}
	}

	// Sort by distance (lower is better), then alphabetically for ties
	slices.SortFunc(matches, func(a, b match) int {
		if a.distance != b.distance {
			if a.distance < b.distance {
				return -1
			}
			return 1
		}
		switch {
		case a.value < b.value:
			return -1
		case a.value > b.value:
			return 1
		default:
			return 0
		}
	})

	// Return top matches
	var results []string
	for i := 0; i < len(matches) && i < maxResults; i++ {
		results = append(results, matches[i].value)
	}

	fuzzyMatchLog.Printf("FindClosestMatches: returning %d match(es) within distance %d", len(results), maxDistance)
	return results
}

// LevenshteinDistance computes the Levenshtein distance between two strings.
// This is the minimum number of single-character edits (insertions, deletions, or substitutions)
// required to change one string into the other.
func LevenshteinDistance(a, b string) int {
	// Optimization: ensure b is the shorter string to minimize row allocation.
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)

	// Early exit for empty strings
	if b == "" {
		return aLen
	}

	// Create a 2D matrix for dynamic programming
	// We only need the previous row, so we can optimize space
	var previousRow, currentRow []int

	// Optimization: use a stack-allocated buffer for small strings to avoid heap allocation.
	var buffer [128]int
	if bLen+1 <= 64 {
		previousRow = buffer[:bLen+1]
		currentRow = buffer[64 : 64+bLen+1]
	} else {
		previousRow = make([]int, bLen+1)
		currentRow = make([]int, bLen+1)
	}

	// Initialize the first row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		previousRow[i] = i
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		currentRow[0] = i // Distance from empty string

		for j := 1; j <= bLen; j++ {
			// Cost of substitution (0 if characters match, 1 otherwise)
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of:
			// - Deletion: previousRow[j] + 1
			// - Insertion: currentRow[j-1] + 1
			// - Substitution: previousRow[j-1] + cost
			deletion := previousRow[j] + 1
			insertion := currentRow[j-1] + 1
			substitution := previousRow[j-1] + cost

			currentRow[j] = min(deletion, min(insertion, substitution))
		}

		// Swap rows for next iteration
		previousRow, currentRow = currentRow, previousRow
	}

	return previousRow[bLen]
}
