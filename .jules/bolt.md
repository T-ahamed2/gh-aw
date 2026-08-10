# Bolt Performance Journal

This journal documents critical performance learnings and patterns discovered during optimizations.

## 2026-08-10 - Optimizing YAML Key Unquoting via Custom Scanner
**Learning:** Regular expressions used for key matching and unquoting at start-of-line boundaries (`(^|\n)([ \t]*)"key":`) are extremely expensive due to regex compilation/runtime and frequent heap allocations. Replacing them with a manual fast-path (`strings.Contains`) and a loop-based scanner utilizing `strings.Index` and `strings.Builder` can reduce processing latency by ~98% (from ~18.5µs down to ~370ns per operation) and heap allocations down to 1 (from 7).
**Action:** When performing simple search-and-replace or string modifications with localized prefix patterns, design custom line/token scanners rather than defaulting to `regexp`. Always verify CPU latency and memory allocations via benchmark suites.
