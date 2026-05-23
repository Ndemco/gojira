# GET /rest/api/3/workflowscheme/project
**operationId:** `getWorkflowSchemeProjectAssociations`
**Summary:** Get workflow scheme project associations

Returns a list of the workflow schemes associated with a list of projects. Each returned workflow scheme includes a list of the requested projects associated with it. Any team-managed or non-existent projects in the request are ignored and no errors are returned.

If the project is associated with the `Default Workflow Scheme` no ID is returned. This is because the way the `Default Workflow Scheme` is stored means it has no ID.

**[Permissions](#permissions) required:** *Administer Jira* [global

## Parameters
- `projectId` [query] (required) []integer(int64) — The ID of a project to return the workflow schemes for. To include multiple projects, provide an ampersand-Jim: onesepar

## Responses
- 200: object:
  - `values` (required): []WorkflowSchemeAssociations
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
