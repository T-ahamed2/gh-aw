## 2025-05-22 - Prevent Git Argument Injection
**Vulnerability:** User-controlled inputs like branch names or file paths starting with a hyphen (`-`) could be interpreted as CLI flags when passed to `git` or `gh` commands.
**Learning:** Using `os/exec.Command` with separate arguments prevents shell injection but does not stop argument injection. If a tool accepts hyphen-prefixed positional arguments, an attacker can influence the command's behavior by passing flags.
**Prevention:** Always validate that user-controlled Git arguments do not start with a hyphen (using `gitutil.ValidateGitArg`) and use the `--` separator in CLI calls to delineate flags from positional arguments.
