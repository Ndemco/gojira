# POST /rest/api/3/permissionscheme/{schemeId}/permission
**operationId:** `createPermissionGrant`
**Summary:** Create permission grant

Creates a permission grant in a permission scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `schemeId` [path] (required) integer(int64) — The ID of the permission scheme in which to create a new permission grant.
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Request Body
Content-Type: `application/json`
object:
  - `holder`: allOf(PermissionHolder)
  - `id`: integer(int64)
  - `permission`: string
  - `self`: string(uri)

## Responses
- 201: object:
  - `holder`: allOf(PermissionHolder)
  - `id`: integer(int64)
  - `permission`: string
  - `self`: string(uri)
- 400: Returned if the value for expand is invalid or the same permission grant is present.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
