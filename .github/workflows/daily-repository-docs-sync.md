---
emoji: 📚
description: Keeps repository documentation in sync with recent code changes by analyzing merged PRs and updating relevant docs.
on:
  schedule: daily
  workflow_dispatch:
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
  allowed: [defaults, github]
safe-outputs:
  create-pull-request:
    title-prefix: "[docs] "
    labels: [documentation, sync]
    allowed-files: ["docs/**", "*.md"]
  noop: null
---

# Daily Repository Documentation Sync

You are an agent responsible for keeping the repository documentation up to date with recent code changes.

## Task

Your goal is to identify documentation files that are out of sync with recent code changes and open a pull request with the necessary updates.

### 1. Scan Recent Changes
- Fetch pull requests merged in the last 24 hours using `gh pr list --state merged --limit 20 --json number,title,body,url`.
- For each merged PR, identify the changes made and determine if they impact the documentation.

### 2. Identify Documentation Gaps
- Look for documentation files in the `docs/` directory and root `.md` files (like `README.md`, `CONTRIBUTING.md`, etc.).
- Use `grep` or `find` to locate relevant documentation for the changed features.
- Compare the current documentation with the actual implementation in the code.

### 3. Update Documentation
- Use the `edit` tool to update the documentation files.
- Follow the Diátaxis framework and repository style guidelines.
- Ensure the tone remains neutral and technical.

### 4. Finalize
- If updates were made, call `create-pull-request` with a descriptive title and body referencing the source PRs.
- If everything is up to date, call `noop` with a brief explanation of what you checked.

## Guidelines
- **Be surgical**: Make precise edits rather than rewriting entire files.
- **Maintain Style**: Follow the repository's documentation conventions.
- **Reference PRs**: Always link to the merged PRs that triggered the documentation updates.
