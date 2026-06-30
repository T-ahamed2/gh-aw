## 2025-05-14 - Regex Character Class Hyphen Placement
**Learning:** In Go's `regexp` syntax (RE2), a hyphen inside a character class (e.g., `[a-z0-9-._]`) is interpreted as a range if it appears between two characters. This can lead to unintended characters being matched (e.g., `9-._` includes characters between '9' and '.').
**Action:** Always place hyphens at the start or end of character classes (e.g., `[a-z0-9._-]`) to ensure they are treated as literal hyphens.
