# Bolt Performance Journal

This journal documents critical performance learnings and optimizations.

## 2023-11-20 - Initial Optimization of Fuzzy Matching
**Learning:** Standard Levenshtein distance dynamic programming allocates 2D matrix rows on the heap for every call, and FindClosestMatches checks all candidates regardless of length, causing significant memory pressure.
**Action:** Implement swapped-string single-row DP with a stack-allocated buffer for short strings to achieve zero heap allocations, and apply length difference short-circuiting to avoid redundant distance calculations.
