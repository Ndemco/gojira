# GET /rest/api/3/config/fieldschemes/{id}/fields
**operationId:** `searchFieldAssociationSchemeFields`
**Summary:** Search field scheme fields

Search for fields belonging to a given field association scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The starting index of the returned fields. Base index: 0.
- `maxResults` [query] integer(int32) — The maximum number of fields to return per page, maximum allowed value is 100.
- `fieldId` [query] []string — The field IDs to filter by, if empty then all fields belonging to a field association scheme will be returned
- `id` [path] (required) integer(int64) — The scheme ID to search for child fields

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []FieldAssociationSchemeFieldSearchResult
- 400: any
- 401: any
- 403: any
- 404: any
