# POST /rest/api/3/worklog/list
**operationId:** `getWorklogsForIds`
**Summary:** Get worklogs

Returns worklog details for a list of worklog IDs.

The returned list of worklogs is limited to 1000 items.

**[Permissions](#permissions) required:** Permission to access Jira, however, worklogs are only returned where either of the following is true:

 *  the worklog is set as *Viewable by All Users*.
 *  the user is a member of a project role or group with permission to view the worklog.

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `prope

## Request Body
Content-Type: `application/json`
object:
  - `ids` (required): []integer(int64)

## Responses
- 200: []object:
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
- 400: Returned if the request contains more than 1000 worklog IDs or is empty.
- 401: Returned if the authentication credentials are incorrect or missing.
