## 2025-05-14 - Git Argument Injection Prevention
**Vulnerability:** Git commands (, , , , ) were executed with user-controlled repository URLs, references, or paths without the `--` separator. This allowed attackers to inject command-line flags by providing inputs starting with a hyphen.
**Learning:** Using `exec.Command` with separate arguments prevents shell injection but does not prevent argument (flag) injection. The `--` separator is necessary in many CLI tools, including Git, to delineate options from positional arguments.
**Prevention:** Always include the `--` separator before positional arguments in `git` and other CLI invocations. For `git checkout`, use `git checkout <ref> --` to disambiguate the reference from paths and mitigate some forms of injection.
## 2025-05-14 - Git Argument Injection Prevention
**Vulnerability:** Git commands (`clone`, `ls-remote`, `archive`, `ls-tree`, `checkout`) were executed with user-controlled repository URLs, references, or paths without the `--` separator. This allowed attackers to inject command-line flags by providing inputs starting with a hyphen.
**Learning:** Using `exec.Command` with separate arguments prevents shell injection but does not prevent argument (flag) injection. The `--` separator is necessary in many CLI tools, including Git, to delineate options from positional arguments.
**Prevention:** Always include the `--` separator before positional arguments in `git` and other CLI invocations. For `git checkout`, use `git checkout <ref> --` to disambiguate the reference from paths and mitigate some forms of injection.
