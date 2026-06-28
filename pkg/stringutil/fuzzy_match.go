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

		// Length check: if absolute difference is more than maxDistance,
		// Levenshtein distance must be at least that difference.
		// This skips the more expensive O(N*M) calculation.
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
	// Ensure b is the shorter string to minimize memory usage for rows
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)

	// Early exit for empty strings
	if b == "" {
		return aLen
	}

	// Use stack-allocated buffers for small strings to avoid heap allocations.
	// 64 is a reasonable limit for typical identifiers and typos.
	var (
		pRowBuf [65]int
		cRowBuf [65]int
		prevRow []int
		currRow []int
	)

	if bLen+1 <= len(pRowBuf) {
		prevRow = pRowBuf[:bLen+1]
		currRow = cRowBuf[:bLen+1]
	} else {
		prevRow = make([]int, bLen+1)
		currRow = make([]int, bLen+1)
	}

	// Initialize the first row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		prevRow[i] = i
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		currRow[0] = i // Distance from empty string

		for j := 1; j <= bLen; j++ {
			// Cost of substitution (0 if characters match, 1 otherwise)
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of:
			// - Deletion: prevRow[j] + 1
			// - Insertion: currRow[j-1] + 1
			// - Substitution: prevRow[j-1] + cost
			deletion := prevRow[j] + 1
			insertion := currRow[j-1] + 1
			substitution := prevRow[j-1] + cost

			currRow[j] = min(deletion, min(insertion, substitution))
		}

		// Swap rows for next iteration
		prevRow, currRow = currRow, prevRow
	}

	return prevRow[bLen]
}
