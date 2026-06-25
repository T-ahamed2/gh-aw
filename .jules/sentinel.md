## 2025-06-25 - [Git Argument Injection]
**Vulnerability:** Argument injection (flag injection) in `git` commands when passing user-controlled strings (like branch names, refs, or repository URLs) as positional arguments.
**Learning:** Even when using `exec.Command` (which avoids shell injection), some commands like `git` can interpret positional arguments starting with `-` as flags, potentially leading to unauthorized configuration changes or other security risks.
**Prevention:** Use the `--` separator to explicitly delineate options from positional arguments in `git` commands (e.g., `git checkout <ref> --`). Always place `--` before any argument that could potentially start with a hyphen.
