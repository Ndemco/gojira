# GET /rest/api/3/universal_avatar/type/{type}/owner/{entityId}
**operationId:** `getAvatars`
**Summary:** Get avatars

Returns the system and custom avatars for a project, issue type or priority.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  for custom project avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the avatar belongs to.
 *  for custom issue type avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for at least one project the issue type is used in.
 *  for syst

## Parameters
- `type` [path] (required) string — The avatar type.
- `entityId` [path] (required) string — The ID of the item the avatar is associated with.

## Responses
- 200: object:
  - `custom`: []Avatar
  - `system`: []Avatar
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the avatar type is invalid, the associated item ID is missing, or the item is not found.
