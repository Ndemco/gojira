# PUT /rest/api/3/comment/{commentId}/properties/{propertyKey}
**operationId:** `setCommentProperty`
**Summary:** Set comment property

Creates or updates the value of a property for a comment. Use this resource to store custom data against a comment.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

**[Permissions](#permissions) required:** either of:

 *  *Edit All Comments* [project permission](https://confluence.atlassian.com/x/yodKLg) to create or update the value of a property on any comment.
 *  *Edit Own Comments* [project p

## Parameters
- `commentId` [path] (required) string — The ID of the comment.
- `propertyKey` [path] (required) string — The key of the property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the comment is not found.
