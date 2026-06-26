## 2026-06-26 - Optimized Fuzzy Matching and Levenshtein Distance
**Learning:** In string-heavy utilities like fuzzy matching, heap allocations for dynamic programming tables can be a significant bottleneck when called frequently in loops.
**Action:** Use stack-allocated buffers for small fixed-size arrays (e.g., [128]int) to eliminate allocations for common cases. Always implement early exit checks based on string length differences to prune expensive (N \cdot M)$ calculations. Ensure the shorter string determines row allocation size in Levenshtein distance.
