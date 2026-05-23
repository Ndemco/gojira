# POST /rest/api/3/issue/bulkfetch
**operationId:** `bulkFetchIssues`
**Summary:** Bulk fetch issues

Returns the details for a set of requested issues. You can request up to 100 issues.

Each issue is identified by its ID or key, however, if the identifier doesn't match an issue, a case-insensitive search and check for moved issues is performed. If a matching issue is found its details are returned, a 302 or other redirect is **not** returned.

Issues will be returned in ascending `id` order. If there are errors, Jira will return a list of issues which couldn't be fetched along with error messa

## Request Body
Content-Type: `application/json`
object:
  - `expand`: []string
  - `fields`: []string
  - `fieldsByKeys`: boolean
  - `issueIdsOrKeys` (required): []string
  - `properties`: []string

## Responses
- 200: object:
  - `issueErrors`: []IssueError
  - `issues`: []IssueBean
- 400: Returned if no issue IDs/keys were present, or more than 100 issue IDs/keys were requested.
- 401: Returned if the authentication credentials are incorrect or missing.
