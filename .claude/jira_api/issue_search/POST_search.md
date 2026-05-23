# POST /rest/api/3/search
**operationId:** `searchForIssuesUsingJqlPost`
**Summary:** Currently being removed. Search for issues using JQL (POST)

Endpoint is currently being removed. [More details](https://developer.atlassian.com/changelog/#CHANGE-2046)

Searches for issues using [JQL](https://confluence.atlassian.com/x/egORLQ).

There is a [GET](#api-rest-api-3-search-get) version of this resource that can be used for smaller JQL query expressions.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Issues are included in the response where the user has:

 *  *Browse projects* [project permission](http

## Request Body
Content-Type: `application/json`
object:
  - `expand`: []string
  - `fields`: []string
  - `fieldsByKeys`: boolean
  - `jql`: string
  - `maxResults`: integer(int32)
  - `properties`: []string
  - `startAt`: integer(int32)
  - `validateQuery`: string

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
