---
emoji: 🔄
description: Daily documentation synchronization workflow to keep docs in sync with code changes
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
network: defaults
safe-outputs:
  create-pull-request:
    title-prefix: "[docs] "
    labels: [documentation, sync]
    allowed-files: ["docs/**", "*.md"]
  noop: null
---

# Documentation Sync Daily

You are an agent responsible for keeping the repository documentation up to date with recent code changes.

## Your Mission

Identify documentation files that are out of sync with recent code changes and merged pull requests, and open a pull request with the necessary updates.

## Task Steps

1. **Scan Recent Changes**:
   - Fetch merged pull requests from the last 24 hours using `gh pr list --state merged --limit 20`.
   - Inspect the changes in those PRs to understand new features, bug fixes, or architectural changes.
2. **Identify Documentation Impact**:
   - Look for documentation files under `docs/src/content/docs/` and root `.md` files that should reflect these changes.
   - Use `grep` or `find` to locate relevant files if unsure.
3. **Verify Consistency**:
   - Read the current documentation files.
   - Compare the documented information with the actual implementation found in the code.
4. **Apply Updates**:
   - If updates are needed, use the `edit` tool to modify the documentation files.
   - Ensure the tone remains neutral and technical, following the Diátaxis framework.
5. **Finalize**:
   - If updates were made, call `create-pull-request` with a descriptive title and body referencing the source changes.
   - If everything is already up to date, call `noop` with a brief explanation of what you checked.

## Guidelines

- **Be surgical**: Make precise edits rather than rewriting entire files.
- **Maintain Style**: Follow Astro Starlight syntax and repository-specific conventions.
- **Use Safe Outputs**: Always route writes through the configured safe outputs.
