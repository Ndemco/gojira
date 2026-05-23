# GET /rest/api/3/issue/{issueIdOrKey}/worklog/{id}
**operationId:** `getWorklog`
**Summary:** Get worklog

Returns a worklog.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, 

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `id` [path] (required) string — The ID of the worklog.
- `expand` [query] string — Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts

## Responses
- 200: object:
  - `author`: allOf(UserDetails)
  - `comment`: any
  - `created`: string(date-time)
  - `id`: string
  - `issueId`: string
  - `properties`: []EntityProperty
  - `self`: string(uri)
  - `started`: string(date-time)
  - `timeSpent`: string
  - `timeSpentSeconds`: integer(int64)
  - `updateAuthor`: allOf(UserDetails)
  - `updated`: string(date-time)
  - `visibility`: allOf(Visibility)
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if:

 *  the issue is not found or the user does not have permission to view it.
 *  the worklog is not found or the user does not have permission to view it.
 *  time tracking is disabled.

.
