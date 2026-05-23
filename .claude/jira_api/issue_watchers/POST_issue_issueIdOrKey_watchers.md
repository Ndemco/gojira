# POST /rest/api/3/issue/{issueIdOrKey}/watchers
**operationId:** `addWatcher`
**Summary:** Add watcher

Adds a user as a watcher of an issue by passing the account ID of the user. For example, `"5b10ac8d82e05b22cc7d4ef5"`. If no user is specified the calling user is added.

This operation requires the **Allow users to watch issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://con

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Request Body
Content-Type: `application/json`
string

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the permission to manage the watcher list.
- 404: Returned if the issue or the user is not found or the user does not have permission to view the issue.
