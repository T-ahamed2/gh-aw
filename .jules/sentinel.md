# Sentinel Security Journal

This file tracks CRITICAL security learnings for this repository.

## 2026-08-09 - Parser-Differential Frontmatter Security Scanner Bypass
**Vulnerability:** Unclosed frontmatter block (`---`) caused the security scanner to treat the file as frontmatter-only and return empty content, completely skipping scanning of malicious payloads that downstream systems or compilers would ultimately compile and execute as markdown body.
**Learning:** Security parsers / pre-processors must fail-secure when parsing malformed or unclosed blocks. Fallbacks should assume the worst-case scenario (that the content is a body payload) to ensure complete security scan coverage.
**Prevention:** Ensure that when a delimiter-based block (like YAML frontmatter) is unclosed or malformed, the parser treats the entire content as the main payload to be scanned instead of treating it as skipped metadata.
