# GET /rest/api/3/avatar/{type}/system
**operationId:** `getAllSystemAvatars`
**Summary:** Get system avatars by type

Returns a list of system avatar details by owner type, where the owner types are issue type, project, user or priority.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `type` [path] (required) string — The avatar type.

## Responses
- 200: object:
  - `system`: []Avatar
- 401: Returned if the authentication credentials are incorrect or missing.
- 500: Returned if an error occurs while retrieving the list of avatars.
