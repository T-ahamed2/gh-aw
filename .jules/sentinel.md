## 2025-05-22 - Flag Injection Protection for Git and GitHub CLI

**Vulnerability:** Argument Injection (Flag Injection) in external command executions. User-controlled strings (like branch names or repository slugs) starting with a hyphen (`-`) could be interpreted as command-line flags by `git` or `gh`, potentially leading to unintended behavior or information disclosure.

**Learning:** While `os/exec` prevents shell injection by not using a shell for execution, it does not prevent the target binary from misinterpreting positional arguments as options if they are prefixed with hyphens.

**Prevention:**
1. Use a validation helper like `gitutil.ValidateGitArg` to explicitly reject hyphen-prefixed arguments for dynamic inputs.
2. Use the `--` separator in commands that support it (like `git clone`, `git ls-remote`, `git ls-tree`, etc.) to delineate options from positional arguments.
