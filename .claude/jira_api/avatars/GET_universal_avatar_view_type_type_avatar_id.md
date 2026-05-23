# GET /rest/api/3/universal_avatar/view/type/{type}/avatar/{id}
**operationId:** `getAvatarImageByID`
**Summary:** Get avatar image by ID

Returns a project, issue type or priority avatar image by ID.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  For system avatars, none.
 *  For custom project avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the avatar belongs to.
 *  For custom issue type avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for at least one project the issue type is used i

## Parameters
- `type` [path] (required) string — The icon type of the avatar.
- `id` [path] (required) integer(int64) — The ID of the avatar.
- `size` [query] string — The size of the avatar image. If not provided the default size is returned.
- `format` [query] string — The format to return the avatar image in. If not provided the original content format is returned.

## Responses
- 200: any
- 200: object
- 200: any
- 200: any
- 400: any
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 400: any
- 400: any
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
