# GET /rest/api/3/dashboard/{id}
**operationId:** `getDashboard`
**Summary:** Get dashboard

Returns a dashboard.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

However, to get a dashboard, the dashboard must be shared with the user or the user must own it. Note, users with the *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) are considered owners of the System dashboard. The System dashboard is considered to be shared with all other users.

## Parameters
- `id` [path] (required) string — The ID of the dashboard.

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
- 404: Returned if the dashboard is not found or the dashboard is not owned by or shared with the user.
