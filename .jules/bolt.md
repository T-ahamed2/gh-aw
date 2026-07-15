## 2025-05-15 - Optimized fuzzy matching with stack allocation and short-circuiting
**Learning:** Levenshtein distance can be optimized to zero heap allocations for common string lengths (<= 64 chars) by using a single-row DP table and a stack-allocated buffer. Swapping inputs ensures the shorter string determines the DP table size. Additionally, a simple length-difference check in the calling function can skip distance calculations entirely for strings that are guaranteed to exceed the maximum allowed distance.
**Action:** Use stack-allocated buffers for DP tables in hot paths when input sizes are typically small, and use inexpensive heuristics to skip expensive algorithms.
## 2025-05-15 - Fixed CI failures caused by missing reports directory
**Learning:** GitHub Actions can fail if a required directory (specified in a tool's configuration) is missing. In this case, the markdown-link-check action failed because the 'reports/' directory did not exist.
**Action:** Always verify that directories referenced in CI workflows exist in the repository, or ensure they are created dynamically if needed. Remove stale/unused CI steps referencing missing directories.
