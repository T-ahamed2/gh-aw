## 2025-05-15 - [Levenshtein Distance Optimization]
**Learning:** Significant performance gains and zero-allocation paths can be achieved in string comparison utilities by using single-row DP tables and stack-allocated buffers for small inputs. Combining this with early-exit length checks in wrapper functions further avoids unnecessary computation for clearly non-matching candidates.
**Action:** Always consider input swapping to minimize DP table size and use stack-allocated buffers (e.g., [65]int) for transient slices in hot loops to eliminate heap pressure.
