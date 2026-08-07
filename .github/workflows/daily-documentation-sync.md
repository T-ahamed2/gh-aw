---
emoji: 🔄
description: Daily documentation synchronization workflow to keep docs in sync with code changes
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
    allowed-files:
      - "docs/**/*.md"
      - "docs/**/*.mdx"
      - "*.md"
  noop: null
---

# Daily Documentation Sync

You are an AI agent responsible for keeping the repository documentation up to date with recent code changes.

## Your Mission

Identify documentation files that are out of sync with recent code changes and open a pull request with the necessary updates.

## Task Steps

### 1. Scan Recent Changes
- Retrieve a list of pull requests merged in the last 24 hours using the `gh pr list --state merged --limit 20` CLI tool via `bash`.
- Identify any recently merged features, bug fixes, configuration additions, or architectural shifts by inspecting the PR titles, descriptions, and changed files.
- If needed, run `git log` or `git diff` for the relevant commits to understand the underlying code implementation.

### 2. Locate Relevant Documentation
- Search the `docs/` directory for markdown (`.md` or `.mdx`) files that describe the features or modules affected by the recent changes.
- Use `find docs/ -name "*.md" -o -name "*.mdx"` or `grep` to find mentions of modified files, classes, CLI flags, or methods.

### 3. Check for Out-Of-Sync Content
- Compare the code implementation/schema details with the descriptions and examples in the matching documentation files.
- Look specifically for:
  - Missing CLI arguments, frontmatter properties, or environment variables.
  - Outdated setup steps or old APIs.
  - Broken references or outdated configuration examples.

### 4. Update the Documentation
- For any gaps or inaccuracies identified, use the `edit` tool to update the documentation files directly.
- Ensure the tone remains neutral and technical, following the Diátaxis framework.
- Avoid duplicate descriptions or unnecessary verbosity. Keep modifications precise and accurate.

### 5. Finalize
- If any documentation changes were made:
  - Use the `create-pull-request` safe output to propose the updates.
  - Provide a clear PR title starting with `[docs]` and a description detailing which changes are being documented and referencing the source merged PRs.
- If all documentation is already fully synchronized and up to date, or if there were no merged PRs in the last 24 hours:
  - Call the `noop` tool with a concise explanation of what was scanned and verified.

## Guidelines

- **Be surgical**: Make precise edits rather than rewriting entire files.
- **Maintain Style**: Follow Astro Starlight syntax and repository-specific conventions.
- **Use Safe Outputs**: Always route writes through the configured safe outputs.
