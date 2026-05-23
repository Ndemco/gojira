# GET /rest/api/3/workflows/capabilities
**operationId:** `workflowCapabilities`
**Summary:** Get available workflow capabilities

Get the list of workflow capabilities for a specific workflow using either the workflow ID, or the project and issue type ID pair. The response includes the scope of the workflow, defined as global/project-based, and a list of project types that the workflow is scoped to. It also includes all rules organised into their broad categories (conditions, validators, actions, triggers, screens) as well as the source location (Atlassian-provided, Connect, Forge).

**[Permissions](#permissions) required:

## Parameters
- `workflowId` [query] string — 
- `projectId` [query] string — 
- `issueTypeId` [query] string — 

## Responses
- 200: object:
  - `connectRules`: []AvailableWorkflowConnectRule
  - `editorScope`: string
  - `forgeRules`: []AvailableWorkflowForgeRule
  - `projectTypes`: []string
  - `systemRules`: []AvailableWorkflowSystemRule
  - `triggerRules`: []AvailableWorkflowTriggers
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
