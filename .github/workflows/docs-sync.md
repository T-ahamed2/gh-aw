---
emoji: 📚
description: Runs daily to identify and update documentation that is out of sync with recent code changes.
on:
  schedule: daily on weekdays
permissions:
  contents: read
  pull-requests: read
  issues: read
  copilot-requests: write
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
safe-outputs:
  create-pull-request:
    allowed-files: ["docs/**", "*.md"]
---

# Documentation Sync

## Task

Your goal is to keep the repository documentation in sync with recent code changes.

1. **Identify Recent Changes**: Use `gh pr list --state merged --limit 50` and `git log` to identify recent changes to the codebase (especially in `pkg/`, `cmd/`, and `internal/`).
2. **Audit Documentation**: Review the existing documentation in the `docs/` directory and root markdown files (e.g., `README.md`, `AGENTS.md`, `CONTRIBUTING.md`, `DEVGUIDE.md`).
3. **Detect Discrepancies**: Look for:
    - New features or CLI commands not yet documented.
    - Changes in logic that make existing documentation inaccurate.
    - Outdated examples, flags, or configuration options.
    - Missing internal documentation for new packages or patterns.
4. **Propose Updates**: If updates are needed, use the `create-pull-request` safe output to propose the changes.
    - Set a descriptive title like `docs: synchronize documentation with recent code changes`.
    - Provide a clear summary of what was updated and why.
    - Limit your changes strictly to documentation files (`docs/**` and `*.md`).
5. **No-op**: If the documentation is already up to date, call `noop` with a brief explanation of your audit.

## Guidelines

- Be concise in your updates.
- Ensure all technical details (commands, flags, paths) are accurate.
- Follow the existing documentation style (Diátaxis framework where applicable).
- If a change is significant, explain the reasoning in the PR body.
