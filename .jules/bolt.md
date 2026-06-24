# Bolt's Journal - Critical Learnings

## 2025-05-14 - Initializing Journal
**Learning:** Initializing the journal for the first time in this codebase.
**Action:** Keep track of performance-related findings.

## 2025-05-14 - Optimized string sanitization
**Learning:** `SanitizeName` was performing multiple passes and compiling regex on every call. Switching to a single-pass `strings.Builder` loop provided a ~9-16x speedup.
**Action:** Always check utility functions for multi-pass patterns or local regex compilation.
