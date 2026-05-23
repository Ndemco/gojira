# POST /rest/api/3/issue/{issueIdOrKey}/worklog
**operationId:** `addWorklog`
**Summary:** Add worklog

Adds a worklog to an issue.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Work on issues* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key the issue.
- `notifyUsers` [query] boolean — Whether users watching the issue are notified by email.
- `adjustEstimate` [query] string — Defines how to update the issue's time estimate, the options are:
- `newEstimate` [query] string — The value to set as the issue's remaining time estimate, as days (\#d), hours (\#h), or minutes (\#m or \#). For example
- `reduceBy` [query] string — The amount to reduce the issue's remaining estimate by, as days (\#d), hours (\#h), or minutes (\#m). For example, *2d*.
- `expand` [query] string — Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts `prop
- `overrideEditableFlag` [query] boolean — Whether the worklog entry should be added to the issue even if the issue is not editable, because jira.issue.editable se

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
- 201: object:
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
 *  `adjustEstimate` is set to `manual` but `reduceBy` is not provided or is invalid.
 *  the user does not have permission to add the worklog.
 *  the request JSON is malformed.
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if the issue is not found or the user does not have permission to view it.
- 413: Returned if the per-issue limit has been breached for one of the following fields:

 *  worklogs
 *  attachments
