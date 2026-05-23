# PUT /rest/api/3/dashboard/bulk/edit
**operationId:** `bulkEditDashboards`
**Summary:** Bulk edit dashboards

Bulk edit dashboards. Maximum number of dashboards to be edited at the same time is 100.

**[Permissions](#permissions) required:** None

The dashboards to be updated must be owned by the user, or the user must be an administrator.

## Request Body
Content-Type: `application/json`
object:
  - `action` (required): string
  - `changeOwnerDetails`: allOf(BulkChangeOwnerDetails)
  - `entityIds` (required): []integer(int64)
  - `extendAdminPermissions`: boolean
  - `permissionDetails`: allOf(PermissionDetails)

## Responses
- 200: object:
  - `action` (required): string
  - `entityErrors`: object
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
