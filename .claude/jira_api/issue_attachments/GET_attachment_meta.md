# GET /rest/api/3/attachment/meta
**operationId:** `getAttachmentMeta`
**Summary:** Get Jira attachment settings

Returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.

Note that there are also [project permissions](https://confluence.atlassian.com/x/yodKLg) that restrict whether users can create and delete attachments.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Responses
- 200: object:
  - `enabled`: boolean
  - `uploadLimit`: integer(int64)
- 401: Returned if the authentication credentials are incorrect or missing.
