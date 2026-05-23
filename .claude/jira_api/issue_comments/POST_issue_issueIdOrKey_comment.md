# POST /rest/api/3/issue/{issueIdOrKey}/comment
**operationId:** `addComment`
**Summary:** Add comment

Adds a comment to an issue.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Add comments* [ project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue containing the comment is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
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
- 201: object:
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
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if the issue is not found or the user does not have permission to view it.
- 413: Returned if the per-issue limit has been breached for one of the following fields:

 *  comments
 *  attachments
