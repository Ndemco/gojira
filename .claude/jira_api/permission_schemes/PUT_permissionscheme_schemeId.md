# PUT /rest/api/3/permissionscheme/{schemeId}
**operationId:** `updatePermissionScheme`
**Summary:** Update permission scheme

Updates a permission scheme. Below are some important things to note when using this resource:

 *  If a permissions list is present in the request, then it is set in the permission scheme, overwriting *all existing* grants.
 *  If you want to update only the name and description, then do not send a permissions list in the request.
 *  Sending an empty list will remove all permission grants from the permission scheme.

If you want to add or delete a permission grant instead of updating the whole

## Parameters
- `schemeId` [path] (required) integer(int64) — The ID of the permission scheme to update.
- `expand` [query] string — Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that p

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `expand`: string
  - `id`: integer(int64)
  - `name` (required): string
  - `permissions`: []PermissionGrant
  - `scope`: allOf(Scope)
  - `self`: string(uri)

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
- 403: Returned if:

 *  the user does not have the necessary permission to update permission schemes.
 *  the Jira instance is Jira Core Free or Jira Software Free. Permission schemes cannot be updated on free plans.
- 404: Returned if the permission scheme is not found.
