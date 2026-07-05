## 2026-07-05 - Git Argument Injection in Remote Fetching
**Vulnerability:** Positional arguments (refs, paths) in external Git commands were susceptible to flag injection because they were not validated for leading hyphens and the `--` separator was missing.
**Learning:** Even when using `exec.Command` (which avoids shell injection), `git` itself may interpret positional arguments starting with `-` as flags (e.g., `--upload-pack`), potentially leading to unauthorized command execution.
**Prevention:** Always use the `--` separator to delineate options from positional arguments in Git commands, and explicitly validate that user-supplied positional arguments do not start with a hyphen. Note that `git checkout` is an exception where `--` changes semantics to path restoration.
