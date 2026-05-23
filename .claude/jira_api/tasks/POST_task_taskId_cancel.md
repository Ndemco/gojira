# POST /rest/api/3/task/{taskId}/cancel
**operationId:** `cancelTask`
**Summary:** Cancel task

Cancels a task.

**[Permissions](#permissions) required:** either of:

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  Creator of the task.

## Parameters
- `taskId` [path] (required) string — The ID of the task.

## Responses
- 202: any
- 400: []string
- 401: []string
- 403: []string
- 404: []string
