## 2026-07-08 - Flag Injection in Git/GitHub CLI Commands
**Vulnerability:** CLI flag/argument injection via dynamic user-controlled inputs (owner, repo, ref, path) passed to `git` or `gh` commands.
**Learning:** Even when using separate arguments in `exec.Command` (preventing shell injection), an attacker can provide values starting with a hyphen (e.g., `-v`, `--upload-pack`) to alter the command's behavior.
**Prevention:** Implement a central validation helper (e.g., `ValidateGitArg`) that explicitly rejects arguments starting with a hyphen and apply it consistently to all dynamic inputs before passing them to CLI tools.
