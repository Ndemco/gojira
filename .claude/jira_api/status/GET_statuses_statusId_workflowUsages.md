# GET /rest/api/3/statuses/{statusId}/workflowUsages
**operationId:** `getWorkflowUsagesForStatus`
**Summary:** Get workflow usages by status

Returns a page of workflows using a given status.

## Parameters
- `statusId` [path] (required) string — The statusId to fetch workflow usages for
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `statusId`: string
  - `workflows`: StatusWorkflowUsagePage
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the status with the given ID does not exist.
