# POST /rest/api/3/filter/{id}/permission
**operationId:** `addSharePermission`
**Summary:** Add share permission

Add a share permissions to a filter. If you add a global share permission (one for all logged-in users or the public) it will overwrite all share permissions for the filter.

Be aware that this operation uses different objects for updating share permissions compared to [Update filter](#api-rest-api-3-filter-id-put).

**[Permissions](#permissions) required:** *Share dashboards and filters* [global permission](https://confluence.atlassian.com/x/x4dKLg) and the user must own the filter.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.

## Request Body
Content-Type: `application/json`
object:
  - `accountId`: string
  - `groupId`: string
  - `groupname`: string
  - `projectId`: string
  - `projectRoleId`: string
  - `rights`: integer(int32)
  - `type` (required): string

## Responses
- 201: []object:
  - `group`: allOf(GroupName)
  - `id`: integer(int64)
  - `project`: allOf(Project)
  - `role`: allOf(ProjectRole)
  - `type` (required): string
  - `user`: allOf(UserBean)
- 400: Returned if:

 *  the request object is invalid. For example, it contains an invalid type, the ID does not match the type, or the project or group is not found.
 *  the user does not own the filter.
 *  the user does not have the required permissions.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the filter is not found.
 *  the user does not have permission to view the filter.
