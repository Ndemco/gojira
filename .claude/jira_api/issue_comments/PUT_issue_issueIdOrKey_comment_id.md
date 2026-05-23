# PUT /rest/api/3/issue/{issueIdOrKey}/comment/{id}
**operationId:** `updateComment`
**Summary:** Update comment

Updates a comment.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue containing the comment is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  *Edit all comments*[ project permission](https://confluence.atlassian.com/x/yodKLg) to update any comm

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `id` [path] (required) string — The ID of the comment.
- `notifyUsers` [query] boolean — Whether users are notified when a comment is updated.
- `overrideEditableFlag` [query] boolean — Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect app users with the 
- `expand` [query] string — Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `rende

## Request Body
Content-Type: `application/json`
object:
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
- 400: Returned if the user does not have permission to edit the comment or the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue or comment is not found or the user does not have permission to view the issue or comment.
