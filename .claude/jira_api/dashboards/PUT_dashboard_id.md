# PUT /rest/api/3/dashboard/{id}
**operationId:** `updateDashboard`
**Summary:** Update dashboard

Updates a dashboard, replacing all the dashboard details with those provided.

**[Permissions](#permissions) required:** None

The dashboard to be updated must be owned by the user.

## Parameters
- `id` [path] (required) string — The ID of the dashboard to update.
- `extendAdminPermissions` [query] boolean — Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](h

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `editPermissions` (required): []SharePermission
  - `name` (required): string
  - `sharePermissions` (required): []SharePermission

## Responses
- 200: object:
  - `automaticRefreshMs`: integer(int32)
  - `description`: string
  - `editPermissions`: []SharePermission
  - `id`: string
  - `isFavourite`: boolean
  - `isWritable`: boolean
  - `name`: string
  - `owner`: allOf(UserBean)
  - `popularity`: integer(int64)
  - `rank`: integer(int32)
  - `self`: string(uri)
  - `sharePermissions`: []SharePermission
  - `systemDashboard`: boolean
  - `view`: string
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
