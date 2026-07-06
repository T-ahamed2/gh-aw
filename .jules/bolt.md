## 2025-06-11 - [Fuzzy Match & String Normalization Optimizations]
**Learning:** `LevenshteinDistance` performance is heavily bottlenecked by heap allocations for the DP table. Using a stack-allocated buffer (`[65]int`) for common small strings (<= 64 bytes) combined with a single-row DP table completely eliminates allocations. Additionally, `strings.Map` is significantly more efficient than manual `strings.Builder` loops for character-level sanitization.
**Action:** Always prefer stack-allocated buffers for small transient slices and use `strings.Map` for single-pass string transformations.
