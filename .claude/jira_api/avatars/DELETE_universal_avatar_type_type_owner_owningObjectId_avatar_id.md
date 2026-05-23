# DELETE /rest/api/3/universal_avatar/type/{type}/owner/{owningObjectId}/avatar/{id}
**operationId:** `deleteAvatar`
**Summary:** Delete avatar

Deletes an avatar from a project, issue type or priority.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `type` [path] (required) string — The avatar type.
- `owningObjectId` [path] (required) string — The ID of the item the avatar is associated with.
- `id` [path] (required) integer(int64) — The ID of the avatar.

## Responses
- 204: Returned if the request is successful.
- 400: Returned if the request is invalid.
- 403: Returned if the user does not have permission to delete the avatar, the avatar is not deletable.
- 404: Returned if the avatar type, associated item ID, or avatar ID is invalid.
