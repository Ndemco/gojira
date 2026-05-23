# GET /rest/api/3/issuetypescheme/project
**operationId:** `getIssueTypeSchemeForProjects`
**Summary:** Get issue type schemes for projects

Returns a [paginated](#pagination) list of issue type schemes and, for each issue type scheme, a list of the projects that use it.

Only issue type schemes used in classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `projectId` [query] (required) []integer(int64) — The list of project IDs. To include multiple project IDs, provide an ampersand-separated list. For example, `projectId=1

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []IssueTypeSchemeProjects
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
