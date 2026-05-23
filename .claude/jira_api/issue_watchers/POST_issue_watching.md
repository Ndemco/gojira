# POST /rest/api/3/issue/watching
**operationId:** `getIsWatchingIssueBulk`
**Summary:** Get is watching issue bulk

Returns, for the user, details of the watched status of issues from a list. If an issue ID is invalid, the returned watched status is `false`.

This operation requires the **Allow users to watch issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yod

## Request Body
Content-Type: `application/json`
object:
  - `issueIds` (required): []string

## Responses
- 200: object:
  - `issuesIsWatching`: object
- 401: Returned if the authentication credentials are incorrect or missing.
