# GET /rest/api/3/search
**operationId:** `searchForIssuesUsingJql`
**Summary:** Currently being removed. Search for issues using JQL (GET)

Endpoint is currently being removed. [More details](https://developer.atlassian.com/changelog/#CHANGE-2046)

Searches for issues using [JQL](https://confluence.atlassian.com/x/egORLQ).

If the JQL query expression is too large to be encoded as a query parameter, use the [POST](#api-rest-api-3-search-post) version of this resource.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Issues are included in the response where the user has:

 *  *Browse projects* 

## Parameters
- `jql` [query] string — The [JQL](https://confluence.atlassian.com/x/egORLQ) that defines the search. Note:
- `startAt` [query] integer(int32) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page. To manage page size, Jira may return fewer items per page where a large 
- `validateQuery` [query] string — Determines how to validate the JQL query and treat the validation results. Supported values are:
- `fields` [query] []string — A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separat
- `expand` [query] string — Use [expand](#expansion) to include additional information about issues in the response. This parameter accepts a comma-
- `properties` [query] []string — A list of issue property keys for issue properties to include in the results. This parameter accepts a comma-separated l
- `fieldsByKeys` [query] boolean — Reference fields by their key (rather than ID).
- `failFast` [query] boolean — Whether to fail the request quickly in case of an error while loading fields for an issue. For `failFast=true`, if one f

## Responses
- 200: object:
  - `expand`: string
  - `issues`: []IssueBean
  - `maxResults`: integer(int32)
  - `names`: object
  - `schema`: object
  - `startAt`: integer(int32)
  - `total`: integer(int32)
  - `warningMessages`: []string
- 400: Returned if the JQL query is invalid.
- 401: Returned if the authentication credentials are incorrect.
