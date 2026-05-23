# GET /rest/api/3/attachment/{id}
**operationId:** `getAttachment`
**Summary:** Get attachment metadata

Returns the metadata for an attachment. Note that the attachment itself is not returned.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.
 *  If attachments are added in private comments, the comme

## Parameters
- `id` [path] (required) string — The ID of the attachment.

## Responses
- 200: object:
  - `author`: allOf(User)
  - `content`: string
  - `created`: string(date-time)
  - `filename`: string
  - `id`: integer(int64)
  - `mimeType`: string
  - `properties`: object
  - `self`: string(uri)
  - `size`: integer(int64)
  - `thumbnail`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if:

 *  the attachment is not found.
 *  attachments are disabled in the Jira settings.
