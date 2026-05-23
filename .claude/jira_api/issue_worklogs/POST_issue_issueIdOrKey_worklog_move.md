# POST /rest/api/3/issue/{issueIdOrKey}/worklog/move
**operationId:** `bulkMoveWorklogs`
**Summary:** Bulk move worklogs

Moves a list of worklogs from one issue to another. This is an experimental API with several limitations:

 *  You can't move more than 5000 worklogs at once.
 *  You can't move worklogs containing an attachment.
 *  You can't move worklogs restricted by project roles.
 *  No notifications will be sent for moved worklogs.
 *  No webhooks or events will be sent for moved worklogs.
 *  No issue history will be recorded for moved worklogs.

Time tracking must be enabled in Jira, otherwise this oper

## Parameters
- `issueIdOrKey` [path] (required) string — 
- `adjustEstimate` [query] string — Defines how to update the issues' time estimate, the options are:
- `overrideEditableFlag` [query] boolean — Whether the work log entry should be moved to and from the issues even if the issues are not editable, because jira.issu

## Request Body
Content-Type: `application/json`
object:
  - `ids`: []integer(int64)
  - `issueIdOrKey`: string

## Responses
- 200: Returned if the request is partially successful.
- 204: Returned if the request is successful.
- 400: Returned if:

 *  `request` is not provided or is invalid
 *  the user does not have permission to move the worklogs
 *  the number of worklogs being moved exceeds the limit
 *  the total size of worklogs being moved is too large
 *  any worklog contains attachments
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if:

 *  the source or destination issue is not found or the user does not have permission to view the issues
 *  at least one of the worklogs is not associated with the provided issue
 *  time tracking is disabled
