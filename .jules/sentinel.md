# Sentinel Security Journal 🛡️

## 2026-06-03 - Subprocess Flag Injection via Git Arguments
**Vulnerability:** External `git` commands (such as `git clone`, `git checkout`, `git ls-remote`, `git archive`, and `git ls-tree`) accept user-controlled positional arguments (such as refs, branches, or file paths) that, if prefixed with a hyphen (`-`), can be interpreted by the shell or subprocess runner as command-line flags. This could lead to flag injection vulnerabilities.
**Learning:** Standard string concatenation or parameter passing in `exec.Command` is insufficient to prevent argument interpretation unless the inputs are explicitly validated or delineated using the `--` end-of-options separator. Certain Git commands (e.g., `git checkout` and `git ls-remote`) do not support `--` for their primary ref/branch argument, meaning strict validation is required at the application level.
**Prevention:** Implement `gitutil.ValidateGitArg` to explicitly reject any hyphen-prefixed strings when constructing arguments for Git commands, and use `--` where supported to robustly partition options from positional path specifications.
