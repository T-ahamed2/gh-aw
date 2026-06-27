## 2026-06-27 - Pre-compiled Regex for Sanitization
**Learning:** Compiling regexes in hot paths like string sanitization is a significant bottleneck. Pre-compiling the possible variants reduced SanitizeName execution time by ~60%.
**Action:** Always check for regexp.MustCompile inside frequently called functions and move them to package-level variables if the pattern is static or has limited variants.
