# GET /rest/api/3/issuesecurityschemes/search
**operationId:** `searchSecuritySchemes`
**Summary:** Search issue security schemes

Returns a [paginated](#pagination) list of issue security schemes.  
If you specify the project ID parameter, the result will contain issue security schemes and related project IDs you filter by. Use \{@link IssueSecuritySchemeResource\#searchProjectsUsingSecuritySchemes(String, String, Set, Set)\} to obtain all projects related to scheme.

Only issue security schemes in the context of classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission]

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of issue security scheme IDs. To include multiple issue security scheme IDs, separate IDs with an ampersand: `i
- `projectId` [query] []string — The list of project IDs. To include multiple project IDs, separate IDs with an ampersand: `projectId=10000&projectId=100

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []SecuritySchemeWithProjects
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
