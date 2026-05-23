# GET /rest/api/3/issue/{issueIdOrKey}/worklog
**operationId:** `getIssueWorklog`
**Summary:** Get issue worklogs

Returns worklogs for an issue (ordered by created time), starting from the oldest worklog or from the worklog started on or after a date and time.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Workloads are only returned where the user has:

 *  *Browse projects* [project permi

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `startedAfter` [query] integer(int64) — The worklog start date and time, as a UNIX timestamp in milliseconds, after which worklogs are returned.
- `startedBefore` [query] integer(int64) — The worklog start date and time, as a UNIX timestamp in milliseconds, before which worklogs are returned.
- `expand` [query] string — Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts`proper

## Responses
- 200: object:
  - `maxResults`: integer(int32)
  - `startAt`: integer(int32)
  - `total`: integer(int32)
  - `worklogs`: []Worklog
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the issue is not found or the user does not have permission to view the issue.
 *  `startAt` or `maxResults` has non-numeric values.
 *  time tracking is disabled.
