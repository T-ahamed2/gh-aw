# Bolt Performance Journal

## 2024-08-03 - Initial Setup
**Learning:** This repository has several Go packages with optimized workflows, custom linters, and helper libraries. To find bottlenecks, we can run `go test -bench` on various packages or search for existing benchmarks.
**Action:** Identify critical paths, check existing benchmark coverage, measure improvements, and document the learnings.

## 2024-08-03 - Optimizing ExtractDomainFromURL via Fast-Path and Logger Guarding
**Learning:** Parsing full URLs with Go's standard library `net/url.Parse` inside high-frequency utility routines introduces significant latency and heap allocation overhead (e.g. ~430ns and 2 allocations per URL) for simple standard paths. By introducing a clean fast-path that scans for common characters (`/`, `:`, `?`, `#`) while checking for complex syntax (`@`, `[`, `]`, `%`) to safely fallback, we can achieve up to 85-87% lower latency with zero allocations. Additionally, guarding internal logger calls with an `.Enabled()` check prevents costly string creation and allocations when debug logging is disabled.
**Action:** Always verify if high-frequency text extraction can employ a zero-allocation manual scanner fast-path, and always wrap active logger calls inside high-frequency utility functions with `.Enabled()` checks.
