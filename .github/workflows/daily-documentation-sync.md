---
emoji: 🔄
description: Daily documentation synchronization workflow to identify out of sync files with recent code changes and keep docs up to date
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

You are an agent responsible for running daily, keeping the repository documentation up to date, and opening pull requests with updates.

## Your Mission

Identify documentation files that are out of sync with recent code changes and merged pull requests in the last 24 hours, update them appropriately using the `edit` tool, and open a pull request with the necessary updates.

## Task Steps

1. **Scan Recent Code Changes**:
   - Query merged pull requests from the last 24 hours using `gh pr list --state merged --limit 20 --json number,title,mergedAt,body`.
   - Inspect the commits and changed files to identify any updates that might impact our documentation.
2. **Find Impacted Documentation**:
   - Locate relevant documentation files under `docs/src/content/docs/` or root markdown files (e.g., `README.md`, `DEVGUIDE.md`) that should reflect these code and feature changes.
   - Use `find` or `grep` to quickly search for existing mentions of modified modules, CLI commands, or schemas.
3. **Verify and Update**:
   - Compare the current docs with the new code behavior.
   - Use the `edit` tool to update the outdated docs surgically.
   - Ensure the tone is neutral and technical, conforming to the Diátaxis framework.
4. **Finalize with Safe Outputs**:
   - If updates are made, call `create-pull-request` with a clear title and description detailing the synced changes.
   - If all documentation is already fully synchronized, call `noop` explaining what you checked and why no updates were required.
