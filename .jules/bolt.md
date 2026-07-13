## 2026-07-13 - [Fuzzy Match Optimization]
**Learning:** In performance-critical search loops (e.g., `FindClosestMatches`), placing the string length-difference short-circuit check before any `strings.ToLower` calls on candidate strings significantly reduces heap allocations and latency for items that fail the length constraint.
**Action:** Always check length constraints or other cheap filters before performing operations that allocate (like `strings.ToLower`, `strings.Split`, or `fmt.Sprintf`).
