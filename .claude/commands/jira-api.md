Look up endpoint schema(s) from the Jira REST API for: $ARGUMENTS

The API reference lives in `.claude/jira_api/` — one directory per resource category (e.g. `issue_search/`, `issues/`, `projects/`), one file per endpoint named `{METHOD}_{path}.md` (e.g. `GET_search_jql.md`).

To find relevant endpoints:
1. Run `ls .claude/jira_api/` to see all resource categories.
2. `ls` the most likely category directory (or directories) to find candidate files.
3. Read the specific endpoint file(s) that match the request.

For each matched endpoint, return the full file contents. Do not implement anything unless the user asks.
