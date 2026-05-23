# POST /rest/api/3/issuetype/{id}/avatar2
**operationId:** `createIssueTypeAvatar`
**Summary:** Load issue type avatar

Loads an avatar for the issue type.

Specify the avatar's local file location in the body of the request. Also, include the following headers:

 *  `X-Atlassian-Token: no-check` To prevent XSRF protection blocking the request, for more information see [Special Headers](#special-request-headers).
 *  `Content-Type: image/image type` Valid image types are JPEG, GIF, or PNG.

For example:  
`curl --request POST \ --user email@example.com:<api_token> \ --header 'X-Atlassian-Token: no-check' \ --head

## Parameters
- `id` [path] (required) string — The ID of the issue type.
- `x` [query] integer(int32) — The X coordinate of the top-left corner of the crop region.
- `y` [query] integer(int32) — The Y coordinate of the top-left corner of the crop region.
- `size` [query] (required) integer(int32) — The length of each side of the crop region.

## Request Body
Content-Type: `*/*`
any

## Responses
- 201: object:
  - `fileName`: string
  - `id` (required): string
  - `isDeletable`: boolean
  - `isSelected`: boolean
  - `isSystemAvatar`: boolean
  - `owner`: string
  - `urls`: object
- 400: Returned if:

 *  an image isn't included in the request.
 *  the image type is unsupported.
 *  the crop parameters extend the crop area beyond the edge of the image.
 *  `cropSize` is missing.
 *  the issue type ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the issue type is not found.
