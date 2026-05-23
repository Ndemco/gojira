# GET /rest/api/3/search/jql
**operationId:** `searchAndReconsileIssuesUsingJql`
**Summary:** Search for issues using JQL enhanced search (GET)

Searches for issues using [JQL](https://confluence.atlassian.com/x/egORLQ). Recent updates might not be immediately visible in the returned search results. If you need [read-after-write](https://developer.atlassian.com/cloud/jira/platform/search-and-reconcile/) consistency, you can utilize the `reconcileIssues` parameter to ensure stronger consistency assurances. This operation can be accessed anonymously.

If the JQL query expression is too large to be encoded as a query parameter, use the [POS

## Parameters
- `jql` [query] string — A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. For performance reasons, this parameter requires a bounde
- `nextPageToken` [query] string — The token for a page to fetch that is not the first page. The first page has a `nextPageToken` of `null`. Use the `nextP
- `maxResults` [query] integer(int32) — The maximum number of items to return per page. To manage page size, API may return fewer items per page where a large n
- `fields` [query] []string — A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separat
- `expand` [query] string — Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority 
- `properties` [query] []string — A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list.
- `fieldsByKeys` [query] boolean — Reference fields by their key (rather than ID). The default is `false`.
- `failFast` [query] boolean — Fail this request early if we can't retrieve all field data.
- `reconcileIssues` [query] []integer(int64) — Strong consistency issue ids to be reconciled with search results. Accepts max 50 ids. This list of ids should be consis

## Responses
- 200: object:
  - `isLast`: boolean
  - `issues`: []IssueBean
  - `names`: object
  - `nextPageToken`: string
  - `schema`: object
- 400: Returned if the search request is invalid
- 401: Returned if the authentication credentials are incorrect or missing.
