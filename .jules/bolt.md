# Bolt Performance Journal

## 2026-07-28 - Regex-based string manipulation in compiler hot paths
**Learning:** Regular expression matches (even when precompiled) add significant CPU and memory allocation overhead in frequently executed compiler routines like YAML unquoting. Replaced regex search-and-replace with a high-performance string scanner utilizing `strings.Index` and `strings.Builder`.
**Action:** Avoid regular expressions inside hot loops or compiler parsing utilities; instead, implement selective fast-paths (`strings.Contains`) and precise loop-based scanners.
