# POST /rest/api/3/comment/list
**operationId:** `getCommentsByIds`
**Summary:** Get comments by IDs

Returns a [paginated](#pagination) list of comments specified by a list of comment IDs.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Comments are returned where the user:

 *  has *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the comment.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If the comme

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts a comm

## Request Body
Content-Type: `application/json`
object:
  - `ids` (required): []integer(int64)

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Comment
- 400: Returned if the request contains more than 1000 IDs or is empty.
