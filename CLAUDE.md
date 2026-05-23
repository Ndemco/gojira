# gojira

A Bubbletea TUI for Jira, written in Go.

## Jira API

When any task involves Jira REST API endpoints — implementing a function, explaining behavior, checking parameters, or anything else — use the `/jira-api` skill to look up the relevant endpoint in `.claude/jira_api_index.md` before proceeding. Do not guess endpoint paths, parameter names, or response shapes.

## Structure

- `jira/client.go` — HTTP client and `newRequest` helper
- `jira/issue.go` — `Issue` type and Bubbletea list.Item interface
- `main.go` / `model.go` / `views.go` / `styles.go` — Bubbletea TUI
