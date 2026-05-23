# POST /rest/api/3/issue/{issueIdOrKey}/changelog/list
**operationId:** `getChangeLogsByIds`
**Summary:** Get changelogs by IDs

Returns changelogs for an issue specified by a list of changelog IDs.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Request Body
Content-Type: `application/json`
object:
  - `changelogIds` (required): []integer(int64)

## Responses
- 200: object:
  - `histories`: []Changelog
  - `maxResults`: integer(int32)
  - `startAt`: integer(int32)
  - `total`: integer(int32)
- 400: Returned if the request is not valid.
- 404: Returned if the issue is not found or the user does not have the necessary permission.
