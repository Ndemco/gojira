# POST /rest/api/3/search/approximate-count
**operationId:** `countIssues`
**Summary:** Count issues using JQL

Provide an estimated count of the issues that match the [JQL](https://confluence.atlassian.com/x/egORLQ). Recent updates might not be immediately visible in the returned output. This endpoint requires JQL to be bounded.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Issues are included in the response where the user has:

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the issue.
 *  If [iss

## Request Body
Content-Type: `application/json`
object:
  - `jql`: string

## Responses
- 200: object:
  - `count`: integer(int64)
- 400: Returned if the JQL query cannot be parsed.
- 401: Returned if the authentication credentials are incorrect.
