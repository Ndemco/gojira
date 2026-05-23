# GET /rest/api/3/permissionscheme/{schemeId}/permission/{permissionId}
**operationId:** `getPermissionSchemeGrant`
**Summary:** Get permission scheme grant

Returns a permission grant.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `schemeId` [path] (required) integer(int64) — The ID of the permission scheme.
- `permissionId` [path] (required) integer(int64) — The ID of the permission grant.
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Responses
- 200: object:
  - `holder`: allOf(PermissionHolder)
  - `id`: integer(int64)
  - `permission`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the permission scheme or permission grant is not found or the user does not have the necessary permission.
