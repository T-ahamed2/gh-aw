# Bolt Performance Journal

## 2026-07-27 - Redundant String Splitting in Logging Paths
**Learning:** Performing multiple `strings.Split` on potentially large strings to extract a line count for logs incurs significant performance and memory allocation overhead. This overhead persists even when debug logging is disabled if not wrapped with conditional checks.
**Action:** Wrap all diagnostic and debug logs in `if logger.Enabled() { ... }` blocks and ensure the string is only split once, reusing the generated slice for downstream processing.
