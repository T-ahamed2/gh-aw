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

		// Performance optimization: skip candidates whose length differs by more
		// than maxDistance. Since each insertion or deletion changes the length by 1,
		// a length difference larger than maxDistance guarantees that the Levenshtein
		// distance will exceed maxDistance. This $O(1)$ check avoids expensive
		// $O(N \times M)$ Levenshtein calculations.
		diff := len(targetLower) - len(candidateLower)
		if diff < -maxDistance || diff > maxDistance {
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
// Performance Optimizations:
//  1. Swaps strings if necessary to make 'b' the shorter string, reducing the DP space.
//  2. Uses a single-row dynamic programming array instead of two rows to minimize space.
//  3. Utilizes a stack-allocated '[65]int' buffer for inputs where the shorter string
//     is <= 64 characters, completely eliminating heap allocations in 100% of typical cases.
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

	// Swap strings if needed so that b is always the shorter string.
	// This reduces the space complexity of the DP table.
	if aLen < bLen {
		a, b = b, a
		aLen, bLen = bLen, aLen
	}

	// Use stack-allocated array for strings <= 64 characters to avoid heap allocations.
	var arr [65]int
	var dp []int
	if bLen+1 <= len(arr) {
		dp = arr[:bLen+1]
	} else {
		dp = make([]int, bLen+1)
	}

	// Initialize the row (distance from empty string)
	for j := 0; j <= bLen; j++ {
		dp[j] = j
	}

	// Calculate distances for each character in string a
	for i := 1; i <= aLen; i++ {
		prevDiagonal := dp[0] // represents dp[i-1][j-1]
		dp[0] = i             // distance from empty string dp[i][0]

		for j := 1; j <= bLen; j++ {
			temp := dp[j] // represents dp[i-1][j]
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}

			// Minimum of Deletion, Insertion, or Substitution
			dp[j] = min(dp[j]+1, min(dp[j-1]+1, prevDiagonal+cost))
			prevDiagonal = temp
		}
	}

	return dp[bLen]
}
