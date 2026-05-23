# GET /rest/api/3/comment/{commentId}/properties
**operationId:** `getCommentPropertyKeys`
**Summary:** Get comment property keys

Returns the keys of all the properties of a comment.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If the comment has visibility restrictions, belongs to the group or has the role visibility is restricted to.

## Parameters
- `commentId` [path] (required) string — The ID of the comment.

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 400: Returned if the comment ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the comment is not found.
