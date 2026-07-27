---
emoji: 🔄
description: Identifies documentation files that are out of sync with recent code changes and opens a pull request with the necessary updates.
on:
  schedule: daily
  workflow_dispatch: null
permissions:
  contents: read
  issues: read
  pull-requests: read
  copilot-requests: write
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
  bash:
    - git
    - find
    - grep
  edit: null
network: defaults
safe-outputs:
  create-pull-request:
    title-prefix: "[docs] "
    labels: [documentation, sync]
    allowed-files: ["docs/**", "*.md"]
  noop: null
---

# Daily Documentation Sync

You are an agent responsible for keeping the repository documentation up to date with recent code changes.

## Your Mission

Scan the repository for recent code changes (merged pull requests and commits in the last 24 hours), identify documentation files that are out of sync, and open a pull request with the necessary updates.

## Task Steps

1. **Scan Recent Changes**:
   - Use the `gh` CLI via Bash to list merged pull requests in the last 24 hours: `gh pr list --state merged --limit 20 --json number,title,mergedAt,body,url`.
   - Review recent commits from the last 24 hours using `git log --since="24 hours ago"`.
   - Inspect the changed files and diffs of those commits and PRs using `git diff` or `git show` to understand new features, bug fixes, CLI changes, or architectural updates.

2. **Identify Documentation Impact**:
   - Locate relevant documentation files under `docs/src/content/docs/` or root markdown files (e.g., `README.md`, `AGENTS.md`, `SKILL.md`) that should reflect these changes.
   - Use `find docs/ -name "*.md"` or `grep` to look for existing files referencing modified features.

3. **Verify Consistency**:
   - Read the current documentation files.
   - Compare the documented information with the actual implementation found in the codebase.
   - Identify any outdated descriptions, missing parameters, incorrect command-line flags, or missing examples.

4. **Apply Updates**:
   - If updates are needed, use the `edit` tool to modify the documentation files precisely.
   - Maintain a neutral, clear, and technical tone, following the Diátaxis framework.
   - Follow Astro Starlight markdown syntax guidelines (such as proper heading levels, tabs, or callouts).

5. **Finalize and Submit**:
   - If documentation updates were made, call the `create-pull-request` safe output to propose a PR. Specify a descriptive branch name, a PR title (formatted as `[docs] Update documentation for recent changes - YYYY-MM-DD`), and a PR description summarizing the changes and references.
   - If all documentation is already completely up to date, call the `noop` safe output with a brief summary explaining what was scanned and why no action was taken.

## Guidelines

- **Be surgical**: Make precise, clean edits instead of rewriting entire files.
- **Maintain Style**: Follow existing repository documentation conventions, standard tone, and formatting.
- **Use Safe Outputs**: Always route writes through the configured safe outputs (`create-pull-request` or `noop`). Never attempt direct git commits or push operations yourself.
