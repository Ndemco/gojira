# POST /rest/api/3/permissionscheme
**operationId:** `createPermissionScheme`
**Summary:** Create permission scheme

Creates a new permission scheme. You can create a permission scheme with or without defining a set of permission grants.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
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
- 201: object:
  - `description`: string
  - `expand`: string
  - `id`: integer(int64)
  - `name` (required): string
  - `permissions`: []PermissionGrant
  - `scope`: allOf(Scope)
  - `self`: string(uri)
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission or the feature is not available in the Jira plan.
