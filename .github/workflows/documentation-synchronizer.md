---
emoji: 📑
description: Keeps repository documentation in sync with recent code changes by identifying gaps and opening pull requests.
on:
  schedule: daily
  workflow_dispatch:
permissions:
  contents: read
  pull-requests: read
  issues: read
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
  cli-proxy: true
network:
  allowed:
    - github
    - defaults
safe-outputs:
  create-pull-request:
    allowed-files: ["docs/**", "*.md"]
    title-prefix: "[docs] "
    labels: ["documentation", "automation"]
  noop: null
---

# Documentation Synchronizer

## Task

Your mission is to keep the repository documentation up to date with the latest code changes.

### 1. Scan Recent Changes
- Use the `github` tool to find pull requests merged in the last 24 hours.
- Identify the key changes, new features, or API updates introduced in these PRs.
- Use `gh pr list --state merged --limit 20` to get a list of recent PRs.
- For relevant PRs, use `gh pr view <number>` to see the description and `gh pr diff <number>` to see the code changes.

### 2. Identify Documentation Gaps
- Explore the `docs/` directory to find relevant documentation files.
- Compare the recent code changes with the existing documentation.
- Look for missing information, outdated examples, or new features that haven't been documented yet.
- Pay attention to CLI command changes, configuration schema updates, and new workflow patterns.

### 3. Update Documentation
- Use the `edit` tool to update the identified documentation files.
- Ensure the documentation is accurate, clear, and follows the project's style guidelines.
- Use GitHub-flavored markdown and maintain consistency with the existing documentation structure.

### 4. Open a Pull Request
- If you have made updates, use `create-pull-request` to submit your changes.
- Provide a clear description of the updates and reference the PRs that triggered them.
- If no documentation updates are needed, call `noop` with a brief explanation.

## Safe Outputs

- Use `create-pull-request` to submit documentation updates.
- Use `noop` when no changes are necessary.
