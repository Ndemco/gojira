# GET /rest/api/3/field/{fieldId}/context
**operationId:** `getContextsForField`
**Summary:** Get custom field contexts

Returns a [paginated](#pagination) list of [ contexts](https://confluence.atlassian.com/adminjiracloud/what-are-custom-field-contexts-991923859.html) for a custom field. Contexts can be returned as follows:

 *  With no other parameters set, all contexts.
 *  By defining `id` only, all contexts from the list of IDs.
 *  By defining `isAnyIssueType`, limit the list of contexts returned to either those that apply to all issue types (true) or those that apply to only a subset of issue types (false)

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `isAnyIssueType` [query] boolean — Whether to return contexts that apply to all issue types.
- `isGlobalContext` [query] boolean — Whether to return contexts that apply to all projects.
- `contextId` [query] []integer(int64) — The list of context IDs. To include multiple contexts, separate IDs with ampersand: `contextId=10000&contextId=10001`.
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
  - `values`: []CustomFieldContext
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
