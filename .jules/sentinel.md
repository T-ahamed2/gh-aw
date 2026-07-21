# Sentinel Security Journal

## 2026-07-21 - Unclosed Frontmatter Security Scanner Bypass
**Vulnerability:** Parser-differential between frontmatter extraction in the markdown security scanner and frontmatter boundary extraction in the compiler/parser. The scanner skipped scanning unclosed frontmatter blocks entirely (assuming they were frontmatter-only), whereas the compiler treated them as having no frontmatter and compiled the entire content as the markdown body. This allowed arbitrary dangerous HTML tags and scripts inside unclosed frontmatter to bypass scanner validation.
**Learning:** Parser differentials occur when two distinct parser implementations analyze the same input and interpret structure differently. One parser was loose (treating unclosed frontmatter as frontmatter) while the other was strict (treating it as plain markdown).
**Prevention:** Always ensure structure parsing and normalization are identical between security scanners and standard compilers. If an opening marker is unclosed, both the scanner and parser must fall back to the same state (body-only text/markdown).
