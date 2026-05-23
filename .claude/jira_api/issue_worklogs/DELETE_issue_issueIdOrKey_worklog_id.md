# DELETE /rest/api/3/issue/{issueIdOrKey}/worklog/{id}
**operationId:** `deleteWorklog`
**Summary:** Delete worklog

Deletes a worklog from an issue.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) i

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `id` [path] (required) string — The ID of the worklog.
- `notifyUsers` [query] boolean — Whether users watching the issue are notified by email.
- `adjustEstimate` [query] string — Defines how to update the issue's time estimate, the options are:
- `newEstimate` [query] string — The value to set as the issue's remaining time estimate, as days (\#d), hours (\#h), or minutes (\#m or \#). For example
- `increaseBy` [query] string — The amount to increase the issue's remaining estimate by, as days (\#d), hours (\#h), or minutes (\#m or \#). For exampl
- `overrideEditableFlag` [query] boolean — Whether the work log entry should be added to the issue even if the issue is not editable, because jira.issue.editable s

## Responses
- 204: Returned if the request is successful.
- 400: Returned if:

 *  `adjustEstimate` is set to `new` but `newEstimate` is not provided or is invalid.
 *  `adjustEstimate` is set to `manual` but `reduceBy` is not provided or is invalid.
 *  the user does not have permission to delete the worklog.
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if:

 *  the issue is not found or user does not have permission to view the issue.
 *  the worklog is not found or the user does not have permission to view it.
 *  time tracking is disabled.
