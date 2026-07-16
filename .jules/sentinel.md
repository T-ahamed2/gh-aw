## 2025-02-11 - Git Flag Injection Protection

**Vulnerability:** Git commands executed via `exec.Command` were vulnerable to "flag injection" because positional arguments like branch names, commit SHAs, or file paths were not validated to ensure they didn't start with a hyphen. An attacker could provide a string like `-u` to be interpreted as a configuration flag by the `git` executable.

**Learning:** While `exec.Command` in Go correctly avoids shell injection by passing arguments separately, it does not inherently prevent the target binary from interpreting those arguments as flags if they start with a hyphen.

**Prevention:** Always validate that untrusted strings passed as positional arguments to external commands do not start with a hyphen. A common pattern is to use the `--` separator to delineate options from positional arguments, but since not all Git commands or versions support this consistently for all argument types, explicit validation of the argument prefix is a robust defense-in-depth measure.
