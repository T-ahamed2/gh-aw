## 2025-05-22 - Git Argument Injection Protection
**Vulnerability:** Flag injection via hyphen-prefixed arguments in `exec.Command`.
**Learning:** Even though `exec.Command` with separate arguments prevents shell injection, Git and other CLIs may interpret arguments starting with a hyphen as flags, potentially leading to arbitrary command execution (e.g., via `--upload-pack` or `--config`).
**Prevention:** Use `ValidateGitArg` to explicitly reject hyphen-prefixed dynamic inputs passed as arguments to external CLIs, or use the `--` separator where supported.
