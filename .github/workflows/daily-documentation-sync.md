---
emoji: 📚
description: Daily documentation sync to identify doc files out of sync with recent code changes and open a PR with updates.
on:
  schedule: daily
  workflow_dispatch: null
permissions:
  contents: read
  pull-requests: read
  issues: read
  copilot-requests: write
tools:
  github:
    mode: gh-proxy
    toolsets: [default]
safe-outputs:
  create-pull-request:
    title-prefix: "[docs] "
    allowed-files: ["docs/**", "*.md"]
  noop: null
---

# Daily Documentation Sync

You are an AI agent responsible for keeping the repository documentation up to date with recent code changes.

## Your Mission

Identify documentation files (under `docs/` and root `.md` files) that are out of sync with recent code changes and merged pull requests from the last 24 hours, and open a pull request with the necessary updates.

## Task Steps

### 1. Scan Recent Changes (Last 24 Hours)

Use the GitHub CLI (`gh`) via Bash or equivalent tools to:
- Retrieve pull requests merged in the last 24 hours:
  ```bash
  gh pr list --state merged --limit 20 --json number,title,body,updatedAt,url
  ```
- Review the specific files changed in those pull requests:
  ```bash
  gh pr diff <PR_NUMBER> --name-only
  ```
- Retrieve recent commits in the last 24 hours if no merged PRs are found.

### 2. Identify Documentation Impact

Compare the code changes against existing documentation under `docs/` and root `.md` files:
- Use `find` or `grep` to locate documentation related to the modified files or components.
- Analyze if any documentation needs to be updated, added, or retired because of the code changes.

### 3. Apply Documentation Updates

If any documentation is out of sync or missing details:
- Use the `edit` tool to update the relevant markdown files.
- Follow a technical, neutral tone (using the Diátaxis framework if applicable).
- Keep changes surgical, precise, and accurate.

### 4. Create Pull Request or No-Op

- **If updates were made**: Call `create-pull-request` with a descriptive title and body outlining the changes. Include references to the merged PRs that prompted the updates.
  - **PR Title Format**: `[docs] Update documentation for features from [date]`
  - **PR Description Formatting**: Use h3 (`###`) or lower for all headers in the description. Wrap detailed change summaries inside `<details><summary>📝 Detailed Changes & References</summary>` tags.
- **If no updates are needed** (all docs are fully synchronized and up to date): Call the `noop` tool with a brief explanation of what you checked and why no updates were necessary.

## Safe Outputs Write Contract

- Create a pull request using `create-pull-request` only when documentation files are modified.
- Always call `noop` if no modifications are made. Do not finish without calling either `create-pull-request` or `noop`.
