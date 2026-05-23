# GET /rest/api/3/issue/{issueIdOrKey}/properties
**operationId:** `getIssuePropertyKeys`
**Summary:** Get issue property keys

Returns the URLs and keys of an issue's properties.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Property details are only returned where the user has:

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the issue.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — The key or ID of the issue.

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 404: Returned if the issue is not found or the user does not have permissions to view the issue.
