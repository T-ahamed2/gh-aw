---
emoji: 🔄
description: Runs daily to keep the repository documentation in sync with recent code changes by opening a pull request
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
network:
  allowed:
    - defaults
    - github
safe-outputs:
  create-pull-request:
    title-prefix: "[docs-sync] "
    labels: [documentation, sync]
    allowed-files: ["docs/**/*.md", "docs/**/*.mdx", "*.md"]
  noop: null
---

# Daily Documentation Synchronizer

You are an agent responsible for keeping the repository documentation up to date with recent code changes and merged pull requests.

## Your Mission

Scan the repository for recent code changes (within the last 24 hours or from recent merged pull requests), identify documentation files that are out of sync with those changes, and open a pull request with the necessary updates.

## Task Steps

### 1. Scan Recent Code Changes

First, gather information about recent commits and merged pull requests.
Use `gh` CLI or git commands to:
- Find all pull requests merged in the last 24 hours:
  ```bash
  gh pr list --state merged --limit 10 --json number,title,mergedAt,body,url
  ```
- Alternatively, check recent commit messages and diffs to identify which features or APIs were added, modified, or removed:
  ```bash
  git log --since="24 hours ago" --oneline
  ```

### 2. Identify Documentation Impact

Analyze the gathered changes to determine which parts of the documentation are affected:
- New features, configuration properties, CLI flags, or schema definitions that are not yet documented.
- Deprecated or removed properties/options that should be removed or updated in existing documentation.
- Review existing Markdown documentation files under `docs/` and root `.md` files. Use `find` or `grep` to locate relevant files:
  ```bash
  find docs/ -name "*.md" -o -name "*.mdx"
  ```

### 3. Verify and Edit Documentation

For any documentation found to be out of sync:
- Use the `edit` tool to update the outdated documentation files surgically.
- Ensure the tone of the documentation remains technical, clear, and objective (following the Diátaxis framework).
- Update examples, API descriptions, and configuration options to match the latest codebase state.
- Keep your changes minimal and precise—do not rewrite entire files unnecessarily.

### 4. Finalize

- **If documentation changes are made**: Open a pull request using the `create-pull-request` safe output with a detailed description of what was synchronized and references to the relevant merged pull requests or commits.
- **If everything is already up to date**: Call `noop` with a brief message explaining what you checked (e.g. recent PRs scanned, and no out-of-sync docs identified).

## Safe Outputs

- Use the configured `create-pull-request` to submit documentation updates.
- Call `noop` with a short explanation when no documentation updates are needed.
