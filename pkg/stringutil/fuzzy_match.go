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

const maxDistance = 3 // Maximum acceptable Levenshtein distance

// FindClosestMatches finds the closest matching strings using Levenshtein distance.
// It returns up to maxResults matches that have a Levenshtein distance of 3 or less.
// Results are sorted by distance (closest first), then alphabetically for ties.
//
// This function is useful for "Did you mean?" suggestions when a user provides
// an unrecognized value (e.g., a typo in an engine name or event type).
//
//nolint:largefunc
func FindClosestMatches(target string, candidates []string, maxResults int) []string {
	fuzzyMatchLog.Printf("FindClosestMatches: target=%q, candidates=%d, maxResults=%d", target, len(candidates), maxResults)

	var matches []match
	targetLower := strings.ToLower(target)
	targetLen := len(targetLower)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		// Skip exact matches
		if targetLower == candidateLower {
			continue
		}

		// Optimization: Levenshtein distance is at least the difference in lengths.
		// If the difference is greater than maxDistance, skip calculation.
		diff := targetLen - len(candidateLower)
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
// This function is optimized to use a single-row dynamic programming array and a stack-allocated
// array for inputs of length <= 64 to completely eliminate heap allocations.
//
//nolint:largefunc
func LevenshteinDistance(a, b string) int {
	aLen := len(a)
	bLen := len(b)

	// Early exit for empty strings
	if a == "" {
		return bLen
	}
	if b == "" {
		return aLen
	}

	// Make b always the shorter string to minimize DP array/slice size
	if aLen < bLen {
		a, b = b, a
		aLen, bLen = bLen, aLen
	}

	// Use stack buffer if short enough to completely avoid heap allocations
	var dp []int
	var stackBuf [65]int
	if bLen+1 <= len(stackBuf) {
		dp = stackBuf[:bLen+1]
	} else {
		dp = make([]int, bLen+1)
	}

	// Initialize the row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		dp[i] = i
	}

	// Calculate distances using a single-row DP table
	for i := 1; i <= aLen; i++ {
		prev := dp[0] // dp[i-1][0]
		dp[0] = i     // dp[i][0]

		for j := 1; j <= bLen; j++ {
			temp := dp[j] // dp[i-1][j], which will become prev for the next column (j+1)

			// Cost of substitution (0 if characters match, 1 otherwise)
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of:
			// - Deletion: dp[j] + 1
			// - Insertion: dp[j-1] + 1
			// - Substitution: prev + cost
			deletion := dp[j] + 1
			insertion := dp[j-1] + 1
			substitution := prev + cost

			dp[j] = min(deletion, min(insertion, substitution))
			prev = temp
		}
	}

	return dp[bLen]
}
