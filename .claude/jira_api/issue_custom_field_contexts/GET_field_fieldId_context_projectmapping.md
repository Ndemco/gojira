# GET /rest/api/3/field/{fieldId}/context/projectmapping
**operationId:** `getProjectContextMapping`
**Summary:** Get project mappings for custom field context

Returns a [paginated](#pagination) list of context to project mappings for a custom field. The result can be filtered by `contextId`. Otherwise, all mappings are returned. Invalid IDs are ignored.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field, for example `customfield\_10000`.
- `contextId` [query] []integer(int64) — The list of context IDs. To include multiple context, separate IDs with ampersand: `contextId=10000&contextId=10001`.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []CustomFieldContextProjectMapping
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
