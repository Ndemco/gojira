# GET /rest/api/3/statuses/{statusId}/projectUsages
**operationId:** `getProjectUsagesForStatus`
**Summary:** Get project usages by status

Returns a page of projects using a given status.

## Parameters
- `statusId` [path] (required) string — The statusId to fetch project usages for
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `projects`: StatusProjectUsagePage
  - `statusId`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the status with the given ID does not exist.
