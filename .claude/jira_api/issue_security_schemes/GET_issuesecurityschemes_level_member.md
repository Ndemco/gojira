# GET /rest/api/3/issuesecurityschemes/level/member
**operationId:** `getSecurityLevelMembers`
**Summary:** Get issue security level members

Returns a [paginated](#pagination) list of issue security level members.

Only issue security level members in the context of classic projects are returned.

Filtering using parameters is inclusive: if you specify both security scheme IDs and level IDs, the result will include all issue security level members from the specified schemes and levels.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of issue security level member IDs. To include multiple issue security level members separate IDs with an amper
- `schemeId` [query] []string — The list of issue security scheme IDs. To include multiple issue security schemes separate IDs with an ampersand: `schem
- `levelId` [query] []string — The list of issue security level IDs. To include multiple issue security levels separate IDs with an ampersand: `levelId
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand opti

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []SecurityLevelMember
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
