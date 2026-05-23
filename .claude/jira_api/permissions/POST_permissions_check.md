# POST /rest/api/3/permissions/check
**operationId:** `getBulkPermissions`
**Summary:** Get bulk permissions

Returns:

 *  for a list of global permissions, the global permissions granted to a user.
 *  for a list of project permissions and lists of projects and issues, for each project permission a list of the projects and issues a user can access or manipulate.

If no account ID is provided, the operation returns details for the logged in user.

Note that:

 *  Invalid project and issue IDs are ignored.
 *  A maximum of 1000 projects and 1000 issues can be checked.
 *  Null values in `globalPermissio

## Request Body
Content-Type: `application/json`
object:
  - `accountId`: string
  - `globalPermissions`: []string
  - `projectPermissions`: []BulkProjectPermissions

## Responses
- 200: object:
  - `globalPermissions` (required): []string
  - `projectPermissions` (required): []BulkProjectPermissionGrants
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
