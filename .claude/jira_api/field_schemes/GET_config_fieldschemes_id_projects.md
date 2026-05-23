# GET /rest/api/3/config/fieldschemes/{id}/projects
**operationId:** `searchFieldAssociationSchemeProjects`
**Summary:** Search field scheme projects

REST Endpoint for searching for projects belonging to a given field association scheme

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The starting index of the returned projects. Base index: 0.
- `maxResults` [query] integer(int32) — The maximum number of projects to return per page, maximum allowed value is 100.
- `projectId` [query] []integer(int64) — The project Ids to filter by, if empty then all projects belonging to a field association scheme will be returned
- `id` [path] (required) integer(int64) — The scheme id to search for associated projects

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []FieldAssociationSchemeProjectSearchResult
- 400: any
- 401: any
- 403: any
- 404: any
