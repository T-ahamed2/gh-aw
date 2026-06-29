## 2025-05-15 - Flag Injection and HTTP Resource Exhaustion Hardening
**Vulnerability:** Positional arguments in `git` commands (owner, repo, ref, path) were not delimited by `--`, and unauthenticated fallback GitHub API calls used `http.DefaultClient` without timeouts.
**Learning:** Even if `exec.Command` uses separate arguments, some tools like `git` may still interpret positional arguments starting with `-` as flags unless `--` is used. Standard library `http.DefaultClient` has no timeout, making the application vulnerable to DoS if a remote server hangs.
**Prevention:** Always use `--` before positional arguments in `git` and similar CLI tools. Explicitly construct `http.Client` with a timeout for all network requests.
