# GET /rest/api/3/field/{fieldId}/context/issuetypemapping
**operationId:** `getIssueTypeMappingsForContexts`
**Summary:** Get issue types for custom field context

Returns a [paginated](#pagination) list of context to issue type mappings for a custom field. Mappings are returned for all contexts or a list of contexts. Mappings are ordered first by context ID and then by issue type ID.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [query] []integer(int64) — The ID of the context. To include multiple contexts, provide an ampersand-separated list. For example, `contextId=10001&
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
  - `values`: []IssueTypeToContextMapping
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
