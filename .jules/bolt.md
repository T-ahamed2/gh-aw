# Bolt Performance Journal

This journal documents critical performance learnings and optimizations made to the `gh-aw` codebase.

## 2025-02-18 - Levenshtein Distance Optimization
**Learning:** Stack allocation for small buffers (such as in `LevenshteinDistance`) completely eliminates heap allocation overhead for typical inputs. Swapping inputs so that the shorter string drives the row size guarantees it fits in stack-allocated buffers up to 64 bytes.
**Action:** Always prioritize stack-allocated or pooled arrays for performance-sensitive slice/buffer allocations where dimensions are bounded.

## 2025-02-18 - Fast-path URL Domain Extraction
**Learning:** Wrapping logging calls with `if logger.Enabled()` blocks prevents costly string formatting and allocation in dry-run/production settings where logging is disabled. Standard HTTP/HTTPS prefixes can be scanned with a zero-allocation fast-path.
**Action:** Check `logger.Enabled()` before any expensive logging argument formatting, and provide direct slice-based fast-paths for highly structured string inputs.

## 2025-02-18 - YAML Key Unquoting & Null Value Cleaning Optimization
**Learning:** Standard regular expressions and broad splitting (e.g. `strings.Split`) inside performance-critical YAML post-processing loops cause massive CPU overhead and high heap allocation rates. Custom scanners using single-pass index searches (`strings.Index`, `strings.IndexByte`) and `strings.Builder` reduce latency and allocations close to zero.
**Action:** Replace regular expressions and split-and-rejoin loops with single-pass index scanners and string builders in hot paths.
