# GET /rest/api/3/field/search
**operationId:** `getFieldsPaginated`
**Summary:** Get fields paginated

Returns a [paginated](#pagination) list of fields for Classic Jira projects. The list can include:

 *  all fields
 *  specific fields, by defining `id`
 *  fields that contain a string in the field name or description, by defining `query`
 *  specific fields that contain a string in the field name or description, by defining `id` and `query`

Use `type` must be set to `custom` to show custom fields only.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `type` [query] []string — The type of fields to search.
- `id` [query] []string — The IDs of the custom fields to return or, where `query` is specified, filter.
- `query` [query] string — String used to perform a case-insensitive partial match with field names or descriptions.
- `orderBy` [query] string — [Order](#ordering) the results by:
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `projectIds` [query] []integer(int64) — The IDs of the projects to filter the fields by. Fields belonging to project Ids that the user does not have access to w

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Field
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
