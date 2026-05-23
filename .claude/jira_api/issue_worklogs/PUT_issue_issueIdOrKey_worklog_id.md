# PUT /rest/api/3/issue/{issueIdOrKey}/worklog/{id}
**operationId:** `updateWorklog`
**Summary:** Update worklog

Updates a worklog.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, 

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key the issue.
- `id` [path] (required) string — The ID of the worklog.
- `notifyUsers` [query] boolean — Whether users watching the issue are notified by email.
- `adjustEstimate` [query] string — Defines how to update the issue's time estimate, the options are:
- `newEstimate` [query] string — The value to set as the issue's remaining time estimate, as days (\#d), hours (\#h), or minutes (\#m or \#). For example
- `expand` [query] string — Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `prope
- `overrideEditableFlag` [query] boolean — Whether the worklog should be added to the issue even if the issue is not editable. For example, because the issue is cl

## Request Body
Content-Type: `application/json`
object:
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
- 400: Returned if:

 *  `adjustEstimate` is set to `new` but `newEstimate` is not provided or is invalid.
 *  the user does not have permission to update the worklog.
 *  the request JSON is malformed.
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if:

 *  the issue is not found or user does not have permission to view the issue.
 *  the worklog is not found or the user does not have permission to view it.
 *  time tracking is disabled.
