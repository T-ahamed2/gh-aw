---
emoji: 🔄
description: Daily documentation synchronization workflow to keep docs in sync with code changes
on:
  schedule: daily
  workflow_dispatch:
permissions:
  contents: read
  pull-requests: read
  issues: read
  copilot-requests: write
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
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

# Documentation Sync

Identify documentation files that are out of sync with recent code changes and merged pull requests, and open a pull request with the necessary updates.

## Task

1. **Scan Recent Changes**:
   - Fetch merged pull requests from the last 24 hours using `gh pr list --state merged --limit 20 --json number,title,mergedAt,body`.
   - Inspect the changes in those PRs to understand new features, bug fixes, or architectural changes.

2. **Identify Documentation Impact**:
   - Look for documentation files under `docs/src/content/docs/` and root `.md` files that should reflect these changes.
   - Read the current documentation files.

3. **Verify Consistency**:
   - Compare the documented information with the actual implementation found in the code.
   - Check if new features are missing or if existing documentation is outdated.

4. **Apply Updates**:
   - If updates are needed, use the `edit` tool to modify the documentation files.
   - Follow the Diátaxis framework: maintain the appropriate type (Tutorial, How-to, Reference, Explanation).
   - Ensure the tone is neutral and technical.

5. **Finalize**:
   - If updates were made, call `create-pull-request` with a descriptive title and body referencing the source changes.
   - If no updates are needed, call `noop` with a brief explanation of what was checked.

## Guidelines

- **Be surgical**: Make precise edits rather than rewriting entire files.
- **Maintain Style**: Follow Astro Starlight syntax and repository-specific conventions.
- **Use Safe Outputs**: Always route writes through the configured safe outputs.
- **No-op**: Explicitly call `noop` if no changes are required after scanning.
