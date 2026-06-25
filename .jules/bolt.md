## 2025-05-14 - String Processing Optimization
**Learning:** High-impact optimizations in this Go codebase often involve replacing multi-pass string manipulations (`strings.Split`, `strings.Join`, `strings.ReplaceAll`) with single-pass `strings.Builder` loops.
**Action:** Always audit string utility functions for multi-pass operations and consider `strings.Builder` with `Grow` for significant allocation reductions.
