# Sentinel Security Journal

## 2026-07-19 - Git Command Flag/Argument Injection
**Vulnerability:** Git commands executed via `exec.Command` on user-supplied references or paths (e.g. `ref`, `path`) can interpret hyphen-prefixed strings as command-line flags (e.g. `--output` or `--help`) rather than positional arguments, potentially leading to command parameter injection or execution bypass.
**Learning:** Even when using Go's `exec.Command` which prevents shell injection by default (as it does not spawn a shell), the target executable itself still parses hyphen-prefixed arguments as options. Therefore, inputs meant for positional arguments must be explicitly validated or separated using a `--` separator.
**Prevention:** Explicitly validate Git arguments using `gitutil.ValidateGitArg` to reject hyphen-prefixed strings before passing them to Git executables. Additionally, when using paths, ensure validation occurs after any path manipulation/normalization (such as `filepath.ToSlash`) to prevent prefix masking.
