# DELETE /rest/api/3/filter/{id}/permission/{permissionId}
**operationId:** `deleteSharePermission`
**Summary:** Delete share permission

Deletes a share permission from a filter.

**[Permissions](#permissions) required:** Permission to access Jira and the user must own the filter.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.
- `permissionId` [path] (required) integer(int64) — The ID of the share permission.

## Responses
- 204: Returned if the request is successful.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the filter is not found.
 *  the user does not own the filter.
