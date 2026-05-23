# GET /rest/api/3/workflow/{workflowId}/project/{projectId}/issueTypeUsages
**operationId:** `getWorkflowProjectIssueTypeUsages`
**Summary:** Get issue types in a project that are using a given workflow

Returns a page of issue types using a given workflow within a project.

## Parameters
- `workflowId` [path] (required) string — The workflow ID
- `projectId` [path] (required) integer(int64) — The project ID
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `issueTypes`: WorkflowProjectIssueTypeUsagePage
  - `projectId`: string
  - `workflowId`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the workflow with the given ID does not exist.
