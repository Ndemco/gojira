# POST /rest/api/3/workflowscheme/{id}/draft/publish
**operationId:** `publishDraftWorkflowScheme`
**Summary:** Publish draft workflow scheme

Publishes a draft workflow scheme.

Where the draft workflow includes new workflow statuses for an issue type, mappings are provided to update issues with the original workflow status to the new workflow status.

This operation is [asynchronous](#async). Follow the `location` link in the response to determine the status of the task and use [Get task](#api-rest-api-3-task-taskId-get) to obtain updates.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://conflu

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme that the draft belongs to.
- `validateOnly` [query] boolean — Whether the request only performs a validation.

## Request Body
Content-Type: `application/json`
object:
  - `statusMappings`: []StatusMapping

## Responses
- 204: Returned if the request is only for validation and is successful.
- 303: object:
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
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: any
