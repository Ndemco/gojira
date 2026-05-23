# POST /rest/api/3/search/jql
**operationId:** `searchAndReconsileIssuesUsingJqlPost`
**Summary:** Search for issues using JQL enhanced search (POST)

Searches for issues using [JQL](https://confluence.atlassian.com/x/egORLQ). Recent updates might not be immediately visible in the returned search results. If you need [read-after-write](https://developer.atlassian.com/cloud/jira/platform/search-and-reconcile/) consistency, you can utilize the `reconcileIssues` parameter to ensure stronger consistency assurances. This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Issues are included in the response where the u

## Request Body
Content-Type: `application/json`
object:
  - `expand`: string
  - `fields`: []string
  - `fieldsByKeys`: boolean
  - `jql`: string
  - `maxResults`: integer(int32)
  - `nextPageToken`: string
  - `properties`: []string
  - `reconcileIssues`: []integer(int64)

## Responses
- 200: object:
  - `isLast`: boolean
  - `issues`: []IssueBean
  - `names`: object
  - `nextPageToken`: string
  - `schema`: object
- 400: Returned if the search request is invalid
- 401: Returned if the authentication credentials are incorrect or missing.
