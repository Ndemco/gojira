# POST /rest/api/3/workflowscheme/update
**operationId:** `updateSchemes`
**Summary:** Update workflow scheme

Updates company-managed and team-managed project workflow schemes. This API doesn't have a concept of draft, so any changes made to a workflow scheme are immediately available. When changing the available statuses for issue types, an [asynchronous task](#async) migrates the issues as defined in the provided mappings.

**[Permissions](#permissions) required:**

 *  *Administer Jira* project permission to update all, including global-scoped, workflow schemes.
 *  *Administer projects* project perm

## Request Body
Content-Type: `application/json`
object:
  - `defaultWorkflowId`: string
  - `description` (required): string
  - `id` (required): string
  - `name` (required): string
  - `statusMappingsByIssueTypeOverride`: []MappingsByIssueTypeOverride
  - `statusMappingsByWorkflows`: []MappingsByWorkflow
  - `version` (required): DocumentVersion
  - `workflowsForIssueTypes`: []WorkflowSchemeAssociation

## Responses
- 200: any
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
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 409: Returned if another workflow configuration update task is ongoing.
