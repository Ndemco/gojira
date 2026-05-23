# GET /rest/api/3/issue/{issueIdOrKey}/comment/{id}
**operationId:** `getComment`
**Summary:** Get comment

Returns a comment.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the comment.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If the comment has visibility restrictions, the user belongs to the group or has the role visibility is restricted to.

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `id` [path] (required) string — The ID of the comment.
- `expand` [query] string — Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `rende

## Responses
- 200: object:
  - `author`: allOf(UserDetails)
  - `body`: any
  - `created`: string(date-time)
  - `id`: string
  - `jsdAuthorCanSeeRequest`: boolean
  - `jsdPublic`: boolean
  - `properties`: []EntityProperty
  - `renderedBody`: string
  - `self`: string
  - `updateAuthor`: allOf(UserDetails)
  - `updated`: string(date-time)
  - `visibility`: allOf(Visibility)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue or comment is not found or the user does not have permission to view the issue or comment.
