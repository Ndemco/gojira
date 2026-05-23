# GET /rest/api/3/comment/{commentId}/properties/{propertyKey}
**operationId:** `getCommentProperty`
**Summary:** Get comment property

Returns the value of a comment property.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If the comment has visibility restrictions, belongs to the group or has the role visibility is restricted to.

## Parameters
- `commentId` [path] (required) string — The ID of the comment.
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 200: object:
  - `key`: string
  - `value`: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the comment or the property is not found.
