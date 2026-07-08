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

	const maxDistance = 3 // Maximum acceptable Levenshtein distance

	var matches []match
	targetLower := strings.ToLower(target)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		// Skip exact matches
		if targetLower == candidateLower {
			continue
		}

		// Early exit: if length difference is greater than maxDistance,
		// the Levenshtein distance must also be greater than maxDistance.
		lenDiff := len(targetLower) - len(candidateLower)
		if lenDiff < 0 {
			lenDiff = -lenDiff
		}
		if lenDiff > maxDistance {
			continue
		}

		distance := LevenshteinDistance(targetLower, candidateLower)

		// Only include if distance is within acceptable range
		if distance <= maxDistance {
			matches = append(matches, match{value: candidate, distance: distance})
		}
	}

	// Sort by distance (lower is better), then alphabetically for ties
	sortMatches(matches)

	// Return top matches
	var results []string
	for i := 0; i < len(matches) && i < maxResults; i++ {
		results = append(results, matches[i].value)
	}

	fuzzyMatchLog.Printf("FindClosestMatches: returning %d match(es) within distance %d", len(results), maxDistance)
	return results
}

type match struct {
	value    string
	distance int
}

func sortMatches(matches []match) {
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
}

// LevenshteinDistance computes the Levenshtein distance between two strings.
// This is the minimum number of single-character edits (insertions, deletions, or substitutions)
// required to change one string into the other.
//
// Bolt Optimization: This implementation uses a single-row dynamic programming table
// to reduce space complexity to O(min(N, M)) and employs a stack-allocated buffer
// for small strings (up to 64 bytes) to eliminate heap allocations in common cases.
func LevenshteinDistance(a, b string) int {
	if a == b {
		return 0
	}

	// Ensure b is the shorter string to minimize space usage
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)

	if b == "" {
		return aLen
	}

	// Use a single row for DP.
	// For small strings, use a stack-allocated buffer to avoid heap allocation.
	var v0 []int
	var stackBuf [65]int
	if bLen < 65 {
		v0 = stackBuf[:bLen+1]
	} else {
		v0 = make([]int, bLen+1)
	}

	// Initialize the row (distance from empty string)
	for i := 0; i <= bLen; i++ {
		v0[i] = i
	}

	for i := 0; i < aLen; i++ {
		prevV0 := v0[0]
		v0[0] = i + 1

		for j := 0; j < bLen; j++ {
			cost := 1
			if a[i] == b[j] {
				cost = 0
			}

			// min(deletion, insertion, substitution)
			// res = min(v0[j+1] + 1, v0[j] + 1, prevV0 + cost)
			res := v0[j+1] + 1
			if v0[j]+1 < res {
				res = v0[j] + 1
			}
			if prevV0+cost < res {
				res = prevV0 + cost
			}

			prevV0 = v0[j+1]
			v0[j+1] = res
		}
	}

	return v0[bLen]
}
