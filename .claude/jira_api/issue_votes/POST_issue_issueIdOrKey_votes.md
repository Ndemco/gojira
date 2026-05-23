# POST /rest/api/3/issue/{issueIdOrKey}/votes
**operationId:** `addVote`
**Summary:** Add vote

Adds the user's vote to an issue. This is the equivalent of the user clicking *Vote* on an issue in Jira.

This operation requires the **Allow users to vote on issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  voting is disabled.
 *  the issue is not found.
