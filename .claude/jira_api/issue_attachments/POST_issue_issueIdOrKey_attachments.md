# POST /rest/api/3/issue/{issueIdOrKey}/attachments
**operationId:** `addAttachment`
**Summary:** Add attachment

Adds one or more attachments to an issue. Attachments are posted as multipart/form-data ([RFC 1867](https://www.ietf.org/rfc/rfc1867.txt)).

Note that:

 *  The request must have a `X-Atlassian-Token: no-check` header, if not it is blocked. See [Special headers](#special-request-headers) for more information.
 *  The name of the multipart/form-data parameter that contains the attachments must be `file`.

The following examples upload a file called *myfile.txt* to the issue *TEST-123*:

#### curl

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue that attachments are added to.

## Request Body
Content-Type: `multipart/form-data`
[]object:
  - `bytes`: []string(byte)
  - `contentType`: string
  - `empty`: boolean
  - `inputStream`: object
  - `name`: string
  - `originalFilename`: string
  - `resource`: Resource
  - `size`: integer(int64)

## Responses
- 200: []object:
  - `author`: allOf(UserDetails)
  - `content`: string
  - `created`: string(date-time)
  - `filename`: string
  - `id`: string
  - `mimeType`: string
  - `self`: string
  - `size`: integer(int64)
  - `thumbnail`: string
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if any of the following is true:

 *  the issue is not found.
 *  the user does not have permission to view the issue.
- 413: Returned if any of the following is true:

 *  the attachments exceed the maximum attachment size for issues.
 *  more than 60 files are requested to be uploaded.
 *  the per-issue limit for attachments has been breached.

See [Configuring file attachments](https://confluence.atlassian.com/x/wIXKM) for details.
