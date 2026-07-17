# Bolt Journal

⚡ Welcome to Bolt's Performance Journal.

## 2026-07-17 - Fuzzy Matching Performance Optimization
**Learning:** In hot search loops like fuzzy matching/typo suggestions (`FindClosestMatches`), calling `strings.ToLower` on candidates and executing full dynamic programming tables for every candidate generates excessive heap allocations and CPU latency. By applying a mathematical length difference short-circuit check (`abs(len(target) - len(candidate)) > maxDistance`) early in the loop, we skip incompatible candidates completely. Furthermore, using a stack-allocated buffer (`[65]int`) for row slices in the Levenshtein Distance DP algorithm eliminates heap allocations for strings up to 64 characters.
**Action:** Always short-circuit string comparison and edit distance loops by comparing length differences first, and use stack-allocated arrays sliced to size for small DP matrices to achieve zero-allocation hot paths.
