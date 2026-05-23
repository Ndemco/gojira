# GET /rest/api/3/issue/{issueIdOrKey}/comment
**operationId:** `getComments`
**Summary:** Get comments

Returns all comments for an issue.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Comments are included in the response where the user has:

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the comment.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If the comment has visibility restrictions, belon

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `orderBy` [query] string — [Order](#ordering) the results by a field. Accepts *created* to sort comments by their created date.
- `expand` [query] string — Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `rende

## Responses
- 200: object:
  - `comments`: []Comment
  - `maxResults`: integer(int32)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
- 400: Returned if `orderBy` is set to a value other than *created*.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue is not found or the user does not have permission to view it.
