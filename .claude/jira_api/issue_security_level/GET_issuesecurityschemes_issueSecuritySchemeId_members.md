# GET /rest/api/3/issuesecurityschemes/{issueSecuritySchemeId}/members
**operationId:** `getIssueSecurityLevelMembers`
**Summary:** Get issue security level members by issue security scheme

Returns issue security level members.

Only issue security level members in context of classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueSecuritySchemeId` [path] (required) integer(int64) — The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) oper
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `issueSecurityLevelId` [query] []string — The list of issue security level IDs. To include multiple issue security levels separate IDs with ampersand: `issueSecur
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand opti

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []IssueSecurityLevelMember
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if no issue security level members are found.
