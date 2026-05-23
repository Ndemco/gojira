# DELETE /rest/api/3/issue/{issueIdOrKey}/worklog
**operationId:** `bulkDeleteWorklogs`
**Summary:** Bulk delete worklogs

Deletes a list of worklogs from an issue. This is an experimental API with limitations:

 *  You can't delete more than 5000 worklogs at once.
 *  No notifications will be sent for deleted worklogs.

Time tracking must be enabled in Jira, otherwise this operation returns an error. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `adjustEstimate` [query] string — Defines how to update the issue's time estimate, the options are:
- `overrideEditableFlag` [query] boolean — Whether the work log entries should be removed to the issue even if the issue is not editable, because jira.issue.editab

## Request Body
Content-Type: `application/json`
object:
  - `ids` (required): []integer(int64)

## Responses
- 200: Returned if the bulk deletion request was partially successful, with a message indicating partial success.
- 204: Returned if the request is successful.
- 400: Returned if:

 *  `request` is not provided or is invalid
 *  the user does not have permission to delete the worklogs
 *  the number of worklogs being deleted exceeds the limit
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if:

 *  the issue is not found or user does not have permission to view the issue
 *  at least one of the worklogs is not associated with the provided issue
 *  time tracking is disabled
