---
emoji: 🔄
description: Daily documentation synchronization workflow to keep docs in sync with recent code changes
on:
  schedule: daily
  workflow_dispatch:
permissions:
  contents: read
  issues: read
  pull-requests: read
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
  cli-proxy: true
network: defaults
safe-outputs:
  create-pull-request:
    title-prefix: "[docs] "
    labels: [documentation, automation]
    allowed-files: ["docs/**", "*.md"]
  noop: null
---

# Daily Repository Documentation Sync

You are an agent responsible for keeping the repository documentation up to date with recent code changes.

## Mission

Scan the repository for recent code changes (merged PRs and commits) from the last 24 hours, identify documentation that needs to be updated to reflect these changes, and open a pull request with the updates.

## Task Steps

1. **Scan Recent Activity**:
   - Use `gh pr list --state merged --limit 20 --json number,title,body,url` to fetch PRs merged in the last 24 hours.
   - Use `gh api /repos/${{ github.repository }}/commits?since=$(date -u -d '24 hours ago' +%Y-%m-%dT%H:%M:%SZ)` to see recent commits.
2. **Analyze Impact**:
   - For each significant change, identify which documentation files are affected.
   - Documentation files are typically found in `docs/` or as root `.md` files (like `README.md`, `CONTRIBUTING.md`, etc.).
3. **Verify and Update**:
   - Read the relevant documentation files.
   - Compare the current documentation with the new implementation.
   - If there's a discrepancy, use the `edit` tool to update the documentation.
4. **Finalize**:
   - If changes were made, call `create-pull-request` with a clear summary of what was updated and why (referencing the source PRs/commits).
   - If no changes are needed, call `noop` with a brief explanation.

## Guidelines

- Follow the project's documentation style and tone.
- Be precise in your updates.
- Ensure all technical details (flags, commands, configuration options) are accurate.
- Use GitHub-flavored markdown.
- Wrap long sections in `<details><summary>...</summary>` in the PR description.
