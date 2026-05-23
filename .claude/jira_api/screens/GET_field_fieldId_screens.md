# GET /rest/api/3/field/{fieldId}/screens
**operationId:** `getScreensForField`
**Summary:** Get screens for a field

Returns a [paginated](#pagination) list of the screens a field is used in.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the field to return screens for.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `expand` [query] string — Use [expand](#expansion) to include additional information about screens in the response. This parameter accepts `tab` w

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ScreenWithTab
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
