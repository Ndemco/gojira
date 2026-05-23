# POST /rest/api/3/workflows/preview
**operationId:** `readWorkflowPreviews`
**Summary:** Preview workflow

Returns a requested workflow within a given project. The response provides a read-only preview of the workflow, omitting full configuration details.

**[Permissions](#permissions) required:**

 *  At least one of the *Administer projects* and *View (read-only) workflow* project permissions

## Request Body
Content-Type: `application/json`
object:
  - `issueTypeIds`: []string
  - `projectId` (required): string
  - `workflowIds`: []string
  - `workflowNames`: []string

## Responses
- 200: object:
  - `statuses`: []JiraWorkflowPreviewStatus
  - `workflows`: []WorkflowPreview
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if one or more previews are not found.
