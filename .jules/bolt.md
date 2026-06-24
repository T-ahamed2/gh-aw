## 2026-06-24 - Efficient String Sanitization
**Learning:** Replacing multi-pass string manipulations (ReplaceAll, Split, Join) and regex-based filtering with a single-pass loop using `strings.Builder` provides significant performance gains (up to 12x faster) and reduces memory pressure by avoiding intermediate allocations.
**Action:** Use single-pass loops and manual character checks in core utility functions that process high volumes of strings or identifiers.

## 2026-06-24 - Avoiding Scope Creep in Performance Tasks
**Learning:** Attempting to fix all CI lint warnings (e.g., `len(s) > 0` vs `s != ""`) and refactoring complex logic (like Docker integrations) while performing a micro-optimization can introduce critical regressions and break established patterns.
**Action:** Keep performance PRs focused on the specific bottleneck identified. Avoid changing function signatures across the package unless absolutely necessary for the optimization.
