# POST /rest/api/3/forge/panel/action/bulk/async
**operationId:** `bulkPinUnpinProjectsAsync`
**Summary:** Bulk pin or unpin issue panel to projects

Bulk pin or unpin an issue panel (added by a Forge app) to or from multiple projects.

The operation runs asynchronously. The response includes a task ID - use the [Get task](#api-rest-api-3-task-taskId-get) endpoint to check progress.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `moduleId` (required): string
  - `projectList` (required): []ProjectPinAction

## Responses
- 202: object:
  - `taskId`: string
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 500: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
