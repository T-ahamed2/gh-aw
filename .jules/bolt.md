## 2026-07-26 - Optimize UnquoteYAMLKey performance
**Learning:** Replacing regex pattern matching and replacement in hot paths (like YAML post-processing or key unquoting) with standard string operations (`strings.Contains`, `strings.Index`, and `strings.Builder`) completely eliminates regex matching/allocation overhead, achieving >98% latency reduction and zero/fewer heap allocations.
**Action:** When performing simple string-level replacements or pattern matches based on line prefixes, favor `strings.Index` scan loops over `regexp.Regexp` replacement functions.
