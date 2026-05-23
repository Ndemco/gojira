# GET /rest/api/3/config/fieldschemes/projects
**operationId:** `getProjectsWithFieldSchemes`
**Summary:** Get projects with field schemes

Get projects with field association schemes. This will be a temporary API but useful when transitioning from the legacy field configuration APIs to the new ones.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The starting index of the returned projects. Base index: 0.
- `maxResults` [query] integer(int32) — The maximum number of projects to return per page, maximum allowed value is 100.
- `projectId` [query] (required) []integer(int64) — List of project ids to filter the results by.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []GetProjectsWithFieldSchemesResponse
- 400: object
- 401: any
- 403: any
- 404: any
