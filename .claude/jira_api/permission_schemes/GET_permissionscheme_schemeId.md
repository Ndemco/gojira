# GET /rest/api/3/permissionscheme/{schemeId}
**operationId:** `getPermissionScheme`
**Summary:** Get permission scheme

Returns a permission scheme.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `schemeId` [path] (required) integer(int64) — The ID of the permission scheme to return.
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Responses
- 200: object:
  - `description`: string
  - `expand`: string
  - `id`: integer(int64)
  - `name` (required): string
  - `permissions`: []PermissionGrant
  - `scope`: allOf(Scope)
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the permission scheme is not found or the user does not have the necessary permission.
