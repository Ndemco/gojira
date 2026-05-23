# GET /rest/api/3/attachment/{id}/expand/human
**operationId:** `expandAttachmentForHumans`
**Summary:** Get all metadata for an expanded attachment

Returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself. For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive. Currently, only the ZIP archive format is supported.

Use this operation to retrieve data that is presented to the user, as this operation returns the metadata for the attachment itself, such as the attachment's ID and name. Otherwise, use 

## Parameters
- `id` [path] (required) string — The ID of the attachment.

## Responses
- 200: object:
  - `entries`: []AttachmentArchiveItemReadable
  - `id`: integer(int64)
  - `mediaType`: string
  - `name`: string
  - `totalEntryCount`: integer(int64)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: The user does not have the necessary permission.
- 404: Returned if:

 *  the attachment is not found.
 *  attachments are disabled in the Jira settings.
- 409: Returned if the attachment is an archive, but not a supported archive format.
