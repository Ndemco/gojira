# POST /rest/api/3/workflowscheme/update/mappings
**operationId:** `getRequiredWorkflowSchemeMappings`
**Summary:** Get required status mappings for workflow scheme update

Gets the required status mappings for the desired changes to a workflow scheme. The results are provided per issue type and workflow. When updating a workflow scheme, status mappings can be provided per issue type, per workflow, or both.

**[Permissions](#permissions) required:**

 *  *Administer Jira* permission to update all, including global-scoped, workflow schemes.
 *  *Administer projects* project permission to update project-scoped workflow schemes.

## Request Body
Content-Type: `application/json`
object:
  - `defaultWorkflowId`: string
  - `id` (required): string
  - `workflowsForIssueTypes` (required): []WorkflowSchemeAssociation

## Responses
- 200: object:
  - `statusMappingsByIssueTypes`: []RequiredMappingByIssueType
  - `statusMappingsByWorkflows`: []RequiredMappingByWorkflows
  - `statuses`: []StatusMetadata
  - `statusesPerWorkflow`: []StatusesPerWorkflow
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
