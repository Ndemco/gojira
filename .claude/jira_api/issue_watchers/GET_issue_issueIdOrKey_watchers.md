# GET /rest/api/3/issue/{issueIdOrKey}/watchers
**operationId:** `getIssueWatchers`
**Summary:** Get issue watchers

Returns the watchers for an issue.

This operation requires the **Allow users to watch issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is ini
 *  If [issue-level 

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Responses
- 200: object:
  - `isWatching`: boolean
  - `self`: string
  - `watchCount`: integer(int32)
  - `watchers`: []UserDetails
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue is not found or the user does not have permission to view it.
