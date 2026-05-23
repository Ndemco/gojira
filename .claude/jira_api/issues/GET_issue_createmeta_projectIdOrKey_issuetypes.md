# GET /rest/api/3/issue/createmeta/{projectIdOrKey}/issuetypes
**operationId:** `getCreateIssueMetaIssueTypes`
**Summary:** Get create metadata issue types for a project

Returns a page of issue type metadata for a specified project. Use the information to populate the requests in [ Create issue](#api-rest-api-3-issue-post) and [Create issues](#api-rest-api-3-issue-bulk-post).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Create issues* [project permission](https://confluence.atlassian.com/x/yodKLg) in the requested projects.

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or key of the project.
- `startAt` [query] integer(int32) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `createMetaIssueType`: []IssueTypeIssueCreateMetadata
  - `issueTypes`: []IssueTypeIssueCreateMetadata
  - `maxResults`: integer(int32)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
