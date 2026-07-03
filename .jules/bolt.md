## 2026-07-03 - Optimize fuzzy matching
**Learning:** Single-row DP and stack-allocation for Levenshtein significantly reduce heap pressure.
**Action:** Always consider shorter string for DP row and use stack buffers for small inputs.
