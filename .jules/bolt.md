## 2026-07-14 - [Optimizing Levenshtein Distance and Fuzzy Matching]
**Learning:** Significant performance gains in Go can be achieved by:
1. Adding a length-difference short-circuit to avoid O(N*M) calculations and allocations (strings.ToLower).
2. Using a stack-allocated buffer (e.g., [65]int) for small inputs to eliminate heap allocations in hot paths.
3. Reducing space complexity to O(min(N, M)) by swapping strings and using a single-row DP approach.
**Action:** Always check for simple heuristics (like length difference) before expensive operations. Use stack buffers for small slices in performance-critical loops.
