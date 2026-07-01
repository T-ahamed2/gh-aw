## 2025-05-15 - Optimizing String Sanitization Utilities
**Learning:** Chained `strings.ReplaceAll` calls are inefficient for multiple replacements as they create multiple intermediate string allocations and perform multiple passes over the data. Similarly, dynamic regex compilation (`regexp.MustCompile`) inside frequently called functions is a major bottleneck. Single-pass loops with pre-allocated `strings.Builder` are significantly faster for character-level transformations.

**Action:** Prefer `strings.Replacer` for static multiple replacements. Use pre-compiled regex patterns for common cases of dynamic sanitization. Always use `sb.Grow(len(input))` when building strings from existing inputs to minimize re-allocations.
