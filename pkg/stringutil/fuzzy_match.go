// Package stringutil provides utility functions for working with strings.
package stringutil

import (
	"slices"
	"strings"

	"github.com/github/gh-aw/pkg/logger"
)

var fuzzyMatchLog = logger.New("stringutil:fuzzy_match")

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

		// Optimization: if length difference is more than maxDistance,
		// they can't be within maxDistance edits.
		diff := len(targetLower) - len(candidateLower)
		if diff < 0 {
			diff = -diff
		}
		if diff > maxDistance {
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
	// Ensure b is the shorter string to minimize space and potentially operations
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)

	// Early exit for empty strings
	if b == "" {
		return aLen
	}

	// Use a single row for DP to save memory.
	// We use a stack-allocated buffer for small strings to avoid heap allocation.
	var row []int
	var buf [65]int
	if bLen+1 <= len(buf) {
		row = buf[:bLen+1]
	} else {
		row = make([]int, bLen+1)
	}

	// Initialize the first row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		row[i] = i
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		lastDiagonal := i - 1
		row[0] = i

		for j := 1; j <= bLen; j++ {
			oldRowJ := row[j]
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of:
			// - Deletion: row[j] + 1
			// - Insertion: row[j-1] + 1
			// - Substitution: lastDiagonal + cost
			row[j] = min(row[j]+1, min(row[j-1]+1, lastDiagonal+cost))
			lastDiagonal = oldRowJ
		}
	}

	return row[bLen]
}
