# Sentinel Journal

## 2026-07-13 - Git Argument Injection Prevention
**Vulnerability:** External Git and GitHub CLI commands were executed with user-supplied arguments (owner, repo, ref, path) without explicitly rejecting hyphen-prefixed strings.
**Learning:** Hyphen-prefixed arguments (e.g., `-v`, `--config`) can be interpreted as flags by command-line tools even when positional arguments are expected, potentially leading to unauthorized configuration changes or information disclosure.
**Prevention:** Always use `gitutil.ValidateGitArg` (or a similar helper) to reject hyphen-prefixed strings before passing them to `exec.Command` or `gh.Exec`. For `git` commands that support it, use the `--` separator to delineate options from positional arguments.
