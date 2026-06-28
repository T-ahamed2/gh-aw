## 2025-05-15 - High-impact string optimization patterns in Go
**Learning:** For performance-critical string utilities like fuzzy matching and sanitization, significant gains can be achieved by:
1. **Stack Allocation:** Use small fixed-size arrays (e.g., `[65]int`) as buffers for slices to avoid heap allocations for common cases (e.g., short identifiers).
2. **Input Swapping:** In algorithms like Levenshtein distance, swapping inputs to ensure the shorter string drives the allocation size minimizes memory usage.
3. **Early Exit:** O(1) checks (like string length difference vs. max distance) can skip expensive O(N*M) calculations.
4. **Regex Pre-compilation:** Storing frequently used regexes in a package-level map avoids the high cost of `regexp.MustCompile` on hot paths.
**Action:** Apply these patterns to all high-frequency string processing functions. Always verify with benchmarks (`go test -bench`).
