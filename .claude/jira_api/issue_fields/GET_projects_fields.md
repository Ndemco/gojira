# GET /rest/api/3/projects/fields
**operationId:** `getProjectFields`
**Summary:** Get fields for projects

Returns a [paginated](#pagination) list of fields for the requested projects and work types.

Only fields that are available for the specified combination of projects and work types are returned. This endpoint allows filtering to specific fields if field IDs are provided.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `projectId` [query] (required) []integer(int64) — The IDs of projects to return fields for.
- `workTypeId` [query] (required) []integer(int64) — The IDs of work types (issue types) to return fields for.
- `fieldId` [query] []string — The IDs of fields to return. If not provided, all fields are returned.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ProjectFieldBean
- 400: Returned if the request parameters are invalid.
- 401: Returned if authentication is missing.
- 403: Returned if the user does not have permission to view the projects or work types.
- 404: Returned if the endpoint is not enabled via feature flag.
