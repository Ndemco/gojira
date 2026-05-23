# GET /rest/api/3/bulk/queue/{taskId}
**operationId:** `getBulkOperationProgress`
**Summary:** Get bulk issue operation progress

Use this to get the progress state for the specified bulk operation `taskId`.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).

If the task is running, this resource will return:

    {"taskId":"10779","status":"RUNNING","progressPercent":65,"submittedBy":{"accountId":"5b10a2844c20165700ede21g"},"created":1690180055963,"started":1690180056206,"updated":169018005829}

If the ta

## Parameters
- `taskId` [path] (required) string — The ID of the task.

## Responses
- 200: object:
  - `created`: string(date-time)
  - `failedAccessibleIssues`: object
  - `invalidOrInaccessibleIssueCount`: integer(int32)
  - `processedAccessibleIssues`: []integer(int64)
  - `progressPercent`: integer(int64)
  - `started`: string(date-time)
  - `status`: string
  - `submittedBy`: User
  - `taskId`: string
  - `totalIssueCount`: integer(int32)
  - `updated`: string(date-time)
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
