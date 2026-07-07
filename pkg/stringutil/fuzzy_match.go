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
	if target == "" {
		return nil
	}
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

		// Optimization: skip Levenshtein calculation if length difference exceeds maxDistance
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
//
// Optimization: This implementation uses a single-row dynamic programming table and
// ensures the shorter string is used for the DP row to minimize space. It also
// employs a stack-allocated buffer for small strings (up to 64 bytes) to eliminate
// heap allocations in the most common cases.
func LevenshteinDistance(a, b string) int {
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)

	// Early exit for empty strings
	if a == "" {
		return bLen
	}
	if b == "" {
		return aLen
	}

	// We only need one row of the matrix for dynamic programming.
	// Use a stack-allocated buffer for small strings to avoid heap allocation.
	var row []int
	var buf [65]int
	if bLen+1 <= len(buf) {
		row = buf[:bLen+1]
	} else {
		row = make([]int, bLen+1)
	}

	// Initialize the row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		row[i] = i
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		prev := i - 1
		row[0] = i

		for j := 1; j <= bLen; j++ {
			// Cost of substitution (0 if characters match, 1 otherwise)
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// substitution: previous value of row[j-1]
			// deletion: current value of row[j]
			// insertion: current value of row[j-1]
			diag := prev
			prev = row[j]

			// Minimum of:
			// - Deletion: row[j] + 1
			// - Insertion: row[j-1] + 1
			// - Substitution: diag + cost
			row[j] = min(row[j]+1, min(row[j-1]+1, diag+cost))
		}
	}

	return row[bLen]
}
