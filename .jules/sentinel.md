## 2026-07-04 - [Git Argument Injection Hardening]
**Vulnerability:** Git command argument injection via user-controlled references or paths starting with a hyphen.
**Learning:** Git positional arguments (URLs, refs, paths) can be interpreted as options if they start with '-', even when passed via `exec.Command`. The `--` separator is essential but `git checkout` specifically needs it BEFORE the reference (`git checkout -- <ref>`) AND benefits from early validation because some versions might still parse the ref as an option if it precedes other flags.
**Prevention:** Always use `--` to delineate options from positional arguments and implement early validation for high-risk inputs like git references.
