# GET /rest/api/3/workflowscheme/{workflowSchemeId}/projectUsages
**operationId:** `getProjectUsagesForWorkflowScheme`
**Summary:** Get projects which are using a given workflow scheme

Returns a page of projects using a given workflow scheme.

## Parameters
- `workflowSchemeId` [path] (required) string — The workflow scheme ID
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `projects`: ProjectUsagePage
  - `workflowSchemeId`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the workflow scheme with the given ID does not exist.
