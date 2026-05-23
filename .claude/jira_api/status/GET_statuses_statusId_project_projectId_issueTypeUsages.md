# GET /rest/api/3/statuses/{statusId}/project/{projectId}/issueTypeUsages
**operationId:** `getProjectIssueTypeUsagesForStatus`
**Summary:** Get issue type usages by status and project

Returns a page of issue types in a project using a given status.

## Parameters
- `statusId` [path] (required) string — The statusId to fetch issue type usages for
- `projectId` [path] (required) string — The projectId to fetch issue type usages for
- `nextPageToken` [query] string — The cursor for pagination
- `maxResults` [query] integer(int32) — The maximum number of results to return. Must be an integer between 1 and 200.

## Responses
- 200: object:
  - `issueTypes`: StatusProjectIssueTypeUsagePage
  - `projectId`: string
  - `statusId`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 404: Returned if the status with the given ID does not exist.
