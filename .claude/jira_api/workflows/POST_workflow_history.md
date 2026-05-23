# POST /rest/api/3/workflow/history
**operationId:** `readWorkflowFromHistory`
**Summary:** Read workflow version from history

Returns a workflow and related statuses for a specified workflow id and version number.

**Note:** Stored workflow data expires after 60 days. Additionally, no data from before the 30th of October 2025 is available.

**[Permissions](#permissions) required:**

 *  *Administer Jira* global permission to access all, including project-scoped, workflows
 *  At least one of the *Administer projects* and *View (read-only) workflow* project permissions to access project-scoped workflows

## Request Body
Content-Type: `application/json`
object:
  - `version`: integer(int64)
  - `workflowId`: string

## Responses
- 200: object:
  - `statuses`: []WorkflowDocumentStatusDTO
  - `workflows`: []WorkflowDocumentDTO
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
