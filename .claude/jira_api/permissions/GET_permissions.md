# GET /rest/api/3/permissions
**operationId:** `getAllPermissions`
**Summary:** Get all permissions

Returns all permissions, including:

 *  global permissions.
 *  project permissions.
 *  global permissions added by plugins.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Responses
- 200: object:
  - `permissions`: object
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
