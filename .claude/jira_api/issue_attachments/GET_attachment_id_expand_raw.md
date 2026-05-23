# GET /rest/api/3/attachment/{id}/expand/raw
**operationId:** `expandAttachmentForMachines`
**Summary:** Get contents metadata for an expanded attachment

Returns the metadata for the contents of an attachment, if it is an archive. For example, if the attachment is a ZIP archive, then information about the files in the archive is returned. Currently, only the ZIP archive format is supported.

Use this operation if you are processing the data without presenting it to the user, as this operation only returns the metadata for the contents of the attachment. Otherwise, to retrieve data to present to the user, use [ Get all metadata for an expanded att

## Parameters
- `id` [path] (required) string — The ID of the attachment.

## Responses
- 200: object:
  - `entries`: []AttachmentArchiveEntry
  - `totalEntryCount`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: The user does not have the necessary permission.
- 404: Returned if:

 *  the attachment is not found.
 *  attachments are disabled in the Jira settings.
- 409: Returned if the attachment is an archive, but not a supported archive format.
