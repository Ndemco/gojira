# GET /rest/api/3/universal_avatar/view/type/{type}
**operationId:** `getAvatarImageByType`
**Summary:** Get avatar image by type

Returns the default project, issue type or priority avatar image.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `type` [path] (required) string — The icon type of the avatar.
- `size` [query] string — The size of the avatar image. If not provided the default size is returned.
- `format` [query] string — The format to return the avatar image in. If not provided the original content format is returned.

## Responses
- 200: any
- 200: object
- 200: any
- 200: any
- 401: any
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: any
- 401: any
- 403: any
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: any
- 403: any
- 404: any
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: any
- 404: any
