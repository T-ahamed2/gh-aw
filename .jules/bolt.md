## 2026-07-31 - [Optimizing YAML Key Unquoting]
**Learning:** Standard regex parsing with `regexp.ReplaceAllString` and `sync.Map` lookup caches incurs severe runtime overhead and heap allocations. Using `strings.Contains` for a zero-allocation fast-path check, followed by a `strings.Builder` and direct character scanning loop, yields a ~98.1% reduction in latency (down to ~364ns/op) and decreases heap allocations from 7 to 1.
**Action:** Always favor direct substring index scanning (`strings.Index`, `strings.Contains`) and `strings.Builder` over regex replacements when matching structured formatting patterns like YAML keys at the start of a line.

## 2026-07-31 - [Repository Feature Validation on Forked Repositories]
**Learning:** Forked repositories often have issues or discussions disabled by default, which can break workflow compilation in CI/CD pipeline runs (`make recompile`) if the compiler validates these as hard failures. Treating missing issues as a warning instead of a hard compilation error (the same as discussions) avoids blocking end-to-end compilation checks.
**Action:** Gracefully handle missing repository features as non-blocking warnings during offline workflow compilation, deferring precise error handling to execution runtime.
