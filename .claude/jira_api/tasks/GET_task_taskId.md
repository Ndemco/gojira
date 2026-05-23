# GET /rest/api/3/task/{taskId}
**operationId:** `getTask`
**Summary:** Get task

Returns the status of a [long-running asynchronous task](#async).

When a task has finished, this operation returns the JSON blob applicable to the task. See the documentation of the operation that created the task for details. Task details are not permanently retained. As of September 2019, details are retained for 14 days although this period may change without notice.

**Deprecation notice:** The required OAuth 2.0 scopes will be updated on June 15, 2024.

 *  `read:jira-work`

**[Permissions

## Parameters
- `taskId` [path] (required) string — The ID of the task.

## Responses
- 200: object:
  - `description`: string
  - `elapsedRuntime` (required): integer(int64)
  - `finished`: integer(int64)
  - `id` (required): string
  - `lastUpdate` (required): integer(int64)
  - `message`: string
  - `progress` (required): integer(int64)
  - `result`: any
  - `self` (required): string(uri)
  - `started`: integer(int64)
  - `status` (required): string
  - `submitted` (required): integer(int64)
  - `submittedBy` (required): integer(int64)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the task is not found.
