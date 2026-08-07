# Bolt's Performance Journal

## 2026-08-07 - Optimizing UnquoteYAMLKey in workflow compilation
**Learning:** Regular-expression-based parsing/replacement (`regexp`) in core utility functions is highly expensive, especially when executed repeatedly. Replacing it with `strings.Contains` for a fast-path bypass and a custom scanner using `strings.Index` and a pre-allocated `strings.Builder` can achieve dramatic latency (~98%) and heap allocation (~85%) reductions.
**Action:** Avoid using `regexp` in performance-critical paths; implement fast-paths via `strings` functions and build outputs using `strings.Builder` with `Grow` pre-allocation.
