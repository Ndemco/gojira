# POST /rest/api/3/field/{fieldId}/context/mapping
**operationId:** `getCustomFieldContextsForProjectsAndIssueTypes`
**Summary:** Get custom field contexts for projects and issue types

Returns a [paginated](#pagination) list of project and issue type mappings and, for each mapping, the ID of a [custom field context](https://confluence.atlassian.com/x/k44fOw) that applies to the project and issue type.

If there is no custom field context assigned to the project then, if present, the custom field context that applies to all projects is returned if it also applies to the issue type or all issue types. If a custom field context is not found, the returned custom field context ID i

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Request Body
Content-Type: `application/json`
object:
  - `mappings` (required): []ProjectIssueTypeMapping

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ContextForProjectAndIssueType
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
