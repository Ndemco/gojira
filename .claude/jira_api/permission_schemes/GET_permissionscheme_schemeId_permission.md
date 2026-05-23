# GET /rest/api/3/permissionscheme/{schemeId}/permission
**operationId:** `getPermissionSchemeGrants`
**Summary:** Get permission scheme grants

Returns all permission grants for a permission scheme.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `schemeId` [path] (required) integer(int64) — The ID of the permission scheme.
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Responses
- 200: object:
  - `expand`: string
  - `permissions`: []PermissionGrant
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the permission schemes is not found or the user does not have the necessary permission.
