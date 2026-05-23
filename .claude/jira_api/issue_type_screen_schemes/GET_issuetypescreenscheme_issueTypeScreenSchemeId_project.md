# GET /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/project
**operationId:** `getProjectsForIssueTypeScreenScheme`
**Summary:** Get issue type screen scheme projects

Returns a [paginated](#pagination) list of projects associated with an issue type screen scheme.

Only company-managed projects associated with an issue type screen scheme are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeScreenSchemeId` [path] (required) integer(int64) — The ID of the issue type screen scheme.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `query` [query] string — 

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ProjectDetails
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
