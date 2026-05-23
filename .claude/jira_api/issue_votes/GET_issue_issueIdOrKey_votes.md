# GET /rest/api/3/issue/{issueIdOrKey}/votes
**operationId:** `getVotes`
**Summary:** Get votes

Returns details about the votes on an issue.

This operation requires the **Allow users to vote on issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is ini
 *  If [

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Responses
- 200: object:
  - `hasVoted`: boolean
  - `self`: string(uri)
  - `voters`: []User
  - `votes`: integer(int64)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  voting is disabled.
 *  the user does not have permission to view the issue.
 *  the issue is not found.
