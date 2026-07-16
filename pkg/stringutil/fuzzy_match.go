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
	if fuzzyMatchLog.Enabled() {
		fuzzyMatchLog.Printf("FindClosestMatches: target=%q, candidates=%d, maxResults=%d", target, len(candidates), maxResults)
	}
	type match struct {
		value    string
		distance int
	}

	const maxDistance = 3 // Maximum acceptable Levenshtein distance

	var matches []match
	targetLower := strings.ToLower(target)
	targetLen := len(targetLower)

	for _, candidate := range candidates {
		// Short-circuit: if length difference is already greater than maxDistance,
		// the Levenshtein distance must be greater than maxDistance.
		candidateLen := len(candidate)
		diff := targetLen - candidateLen
		if diff < 0 {
			diff = -diff
		}
		if diff > maxDistance {
			continue
		}

		candidateLower := strings.ToLower(candidate)

		// Skip exact matches
		if targetLower == candidateLower {
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

	if fuzzyMatchLog.Enabled() {
		fuzzyMatchLog.Printf("FindClosestMatches: returning %d match(es) within distance %d", len(results), maxDistance)
	}
	return results
}

// LevenshteinDistance computes the Levenshtein distance between two strings.
// This is the minimum number of single-character edits (insertions, deletions, or substitutions)
// required to change one string into the other.
func LevenshteinDistance(a, b string) int {
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}
	if a == b {
		return 0
	}

	// Ensure b is the shorter string to minimize space complexity
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen, bLen := len(a), len(b)

	// Use a stack-allocated buffer for small strings to avoid heap allocation.
	// 65 is chosen as a reasonable limit for typical workflow identifiers and engine names.
	var row []int
	var buffer [65]int
	if bLen+1 <= len(buffer) {
		row = buffer[:bLen+1]
	} else {
		row = make([]int, bLen+1)
	}

	// Initialize the row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		row[i] = i
	}

	// Calculate distances using single-row DP optimization
	for i := 1; i <= aLen; i++ {
		prevRowCell := i - 1 // Stores value of (i-1, j-1)
		row[0] = i           // Distance from empty string (i, 0)

		for j := 1; j <= bLen; j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of:
			// - Deletion: row[j] + 1 (from previous row i-1)
			// - Insertion: row[j-1] + 1 (from current row i)
			// - Substitution: prevRowCell + cost (from cell i-1, j-1)
			deletion := row[j] + 1
			insertion := row[j-1] + 1
			substitution := prevRowCell + cost

			nextPrevRowCell := row[j] // Save (i-1, j) for next j's substitution
			row[j] = min(deletion, min(insertion, substitution))
			prevRowCell = nextPrevRowCell
		}
	}

	return row[bLen]
}
