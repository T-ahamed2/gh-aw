## 2025-05-15 - Optimizing String Sanitization Utilities
**Learning:** Chained `strings.ReplaceAll` calls are inefficient for multiple replacements as they create multiple intermediate string allocations and perform multiple passes over the data. Similarly, dynamic regex compilation (`regexp.MustCompile`) inside frequently called functions is a major bottleneck. Single-pass loops with pre-allocated `strings.Builder` are significantly faster for character-level transformations.

**Action:** Prefer `strings.Replacer` for static multiple replacements. Use pre-compiled regex patterns for common cases of dynamic sanitization. Always use `sb.Grow(len(input))` when building strings from existing inputs to minimize re-allocations.
## 2025-05-16 - Go Idiomatic Length Checks and Regex Ranges
**Learning:** The custom project linters enforce `s != ""` over `len(s) > 0` for string length checks. Additionally, regex character classes like `[0-9-_]` are dangerous in Go as they interpret the hyphen as a range between the preceding and following character (e.g., `9` to `_`), leading to unintended matches or runtime panics if the range is invalid (e.g., start > end).

**Action:** Always place hyphens at the start or end of character classes (`[0-9_-]`) to avoid range interpretation. Prefer `s != ""` for string emptiness checks to satisfy project linting rules.
