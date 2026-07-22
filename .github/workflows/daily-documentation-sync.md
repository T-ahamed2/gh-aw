---
emoji: 🔄
description: Daily documentation synchronization workflow to identify out-of-sync docs and keep them updated with recent code changes via pull request.
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

Identify documentation files that are out of sync with recent code changes and merged pull requests, and open a pull request with the necessary updates.

## Task Steps

### 1. Scan Recent Changes
- Fetch merged pull requests from the last 24 hours using `gh pr list --state merged --limit 20 --json number,title,mergedAt,body,url`.
- Inspect the changes in those PRs to understand new features, bug fixes, or architectural changes.

### 2. Identify Documentation Impact
- Locate documentation files under `docs/src/content/docs/` and root `.md` files that should reflect these changes.
- Use `find` or `grep` to search for relevant sections/topics in the documentation.

### 3. Verify Consistency
- Read the current documentation files.
- Compare the documented information with the actual implementation found in the code.

### 4. Apply Updates
- If updates are needed, use the `edit` tool to modify the documentation files.
- Ensure the tone remains neutral and technical, following the Diátaxis framework and Astro Starlight syntax conventions.
- Be surgical: Make precise, focused edits rather than rewriting entire files.

### 5. Finalize
- If documentation updates were made, call `create-pull-request` with a descriptive title and body referencing the source changes (e.g. `[docs] Update documentation for features from [date]`).
- If everything is already up to date, call `noop` with a brief, clear explanation of what you checked.

## Safe Outputs

- Use `create-pull-request` to propose documentation updates.
- Use `noop` with a short explanation when no documentation updates are needed.
