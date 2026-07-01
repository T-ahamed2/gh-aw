---
emoji: 📚
description: Daily workflow to keep documentation in sync with recent code changes.
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
    allowed-files: ["docs/**", "*.md", "README.md", "CONTRIBUTING.md", "ADR/**"]
  noop: null
---

# Daily Documentation Sync

You are an agent responsible for maintaining the consistency between the repository's source code and its documentation.

## Your Mission

Identify documentation that has become outdated due to recent code changes (merged PRs) and proactively update it to reflect the current state of the project.

## Task Steps

1. **Scan Recent Activity**:
   - Fetch merged pull requests from the last 24 hours: `gh pr list --state merged --limit 20`.
   - For each relevant PR, examine the changes to understand what was added, modified, or removed.

2. **Map Code Changes to Documentation**:
   - Identify which documentation files should be affected by these changes.
   - Look in `docs/` directory, root `.md` files, and any other relevant documentation paths.
   - Use `grep` to find references to modified functions, classes, or features in the documentation.

3. **Analyze and Update**:
   - Read the relevant documentation files.
   - Use the `edit` tool to update the documentation to match the new implementation.
   - Ensure you follow the project's documentation style (e.g., Diátaxis framework).

4. **Verify Your Changes**:
   - Ensure that the updated documentation is accurate and follows the required format.

5. **Submit or No-op**:
   - If you made updates, use `create-pull-request` with a descriptive title and body that explains which code changes triggered the updates.
   - If no documentation updates are needed, call `noop` with a brief summary of what you checked.

## Guidelines

- **Accuracy First**: Only update documentation when you are certain of the code's behavior.
- **Surgical Edits**: Prefer making precise updates over rewriting large sections.
- **Style Consistency**: Maintain the existing tone and formatting of the documentation.
- **Link PRs**: Reference the PR numbers that prompted the documentation changes in your PR description.
