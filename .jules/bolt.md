## 2024-11-20 - Levenshtein Distance Optimization in pkg/stringutil
**Learning:** Swapping string parameters (when `len(a) < len(b)`) inside the fuzzy match algorithm allows allocating the dynamic programming row of size `len(a)+1` on a stack-allocated buffer (e.g., `[65]int`) for common input lengths (<= 64 chars). Length short-circuiting in `FindClosestMatches` also drastically limits execution frequency.
**Action:** Use fixed-size array stack buffers as fallback structures for slice allocations in hot loops when input sizes have typical upper bounds.

## 2026-07-25 - Extract Domain Fast-Path & Logging Guardrails
**Learning:** Unconditional structured/namespace logging calls (e.g., `urlsLog.Printf`) evaluate arguments and invoke formatting routines even when logging is disabled, introducing ~70ns of overhead. Additionally, standard library `net/url.Parse` is extremely heavy and allocates multiple structs, taking ~450ns for standard HTTP/HTTPS URLs.
**Action:** Always guard logger calls with `.Enabled()` checks in hot execution paths, and implement zero-allocation string slicing fast-paths to bypass full URL parsing for standard schema patterns.
