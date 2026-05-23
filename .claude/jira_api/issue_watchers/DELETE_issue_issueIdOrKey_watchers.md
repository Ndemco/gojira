# DELETE /rest/api/3/issue/{issueIdOrKey}/watchers
**operationId:** `removeWatcher`
**Summary:** Delete watcher

Deletes a user as a watcher of an issue.

This operation requires the **Allow users to watch issues** option to be *ON*. This option is set in General configuration for Jira. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0

## Responses
- 204: Returned if the request is successful.
- 400: Returned if `accountId` is not supplied.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the permission to manage the watcher list.
- 404: Returned if the issue or the user is not found or the user does not have permission to view the issue.
