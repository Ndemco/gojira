# GET /rest/api/3/filter/{id}/permission/{permissionId}
**operationId:** `getSharePermission`
**Summary:** Get share permission

Returns a share permission for a filter. A filter can be shared with groups, projects, all logged-in users, or the public. Sharing with all logged-in users or the public is known as a global share permission.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None, however, a share permission is only returned for:

 *  filters owned by the user.
 *  filters shared with a group that the user is a member of.
 *  filters shared with a private project that the us

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.
- `permissionId` [path] (required) integer(int64) — The ID of the share permission.

## Responses
- 200: object:
  - `group`: allOf(GroupName)
  - `id`: integer(int64)
  - `project`: allOf(Project)
  - `role`: allOf(ProjectRole)
  - `type` (required): string
  - `user`: allOf(UserBean)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the filter is not found.
 *  the permission is not found.
 *  the user does not have permission to view the filter.
