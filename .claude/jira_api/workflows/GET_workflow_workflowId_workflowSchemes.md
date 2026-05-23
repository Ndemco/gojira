# GET /rest/api/3/workflow/{workflowId}/workflowSchemes
**operationId:** `getWorkflowSchemeUsagesForWorkflow`
**Summary:** Get workflow schemes which are using a given workflow

Returns a page of workflow schemes using a given workflow.

## Parameters
- `workflowId` [path] (required) string — The workflow ID
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `workflowId`: string
  - `workflowSchemes`: WorkflowSchemeUsagePage
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the workflow with the given ID does not exist.
