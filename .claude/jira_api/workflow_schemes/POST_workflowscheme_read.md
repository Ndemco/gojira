# POST /rest/api/3/workflowscheme/read
**operationId:** `readWorkflowSchemes`
**Summary:** Bulk get workflow schemes

Returns a list of workflow schemes by providing workflow scheme IDs or project IDs.

**[Permissions](#permissions) required:**

 *  *Administer Jira* global permission to access all, including project-scoped, workflow schemes
 *  *Administer projects* project permissions to access project-scoped workflow schemes

## Request Body
Content-Type: `application/json`
object:
  - `projectIds`: []string
  - `workflowSchemeIds`: []string

## Responses
- 200: []object:
  - `defaultWorkflow`: WorkflowMetadataRestModel
  - `description`: string
  - `id` (required): string
  - `name` (required): string
  - `scope` (required): WorkflowScope
  - `taskId`: string
  - `version` (required): DocumentVersion
  - `workflowsForIssueTypes` (required): []WorkflowMetadataAndIssueTypeRestModel
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
