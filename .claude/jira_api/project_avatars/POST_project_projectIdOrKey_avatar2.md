# POST /rest/api/3/project/{projectIdOrKey}/avatar2
**operationId:** `createProjectAvatar`
**Summary:** Load project avatar

Loads an avatar for a project.

Specify the avatar's local file location in the body of the request. Also, include the following headers:

 *  `X-Atlassian-Token: no-check` To prevent XSRF protection blocking the request, for more information see [Special Headers](#special-request-headers).
 *  `Content-Type: image/image type` Valid image types are JPEG, GIF, or PNG.

For example:  
`curl --request POST `

`--user email@example.com:<api_token> `

`--header 'X-Atlassian-Token: no-check' `

`--hea

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or (case-sensitive) key of the project.
- `x` [query] integer(int32) — The X coordinate of the top-left corner of the crop region.
- `y` [query] integer(int32) — The Y coordinate of the top-left corner of the crop region.
- `size` [query] integer(int32) — The length of each side of the crop region.

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
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to administer the project or an anonymous call is made to the operation.
- 404: Returned if the project is not found or the user does not have permission to view the project.
