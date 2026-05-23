# PUT /rest/api/3/project/{projectKeyOrId}/permissionscheme
**operationId:** `assignPermissionScheme`
**Summary:** Assign permission scheme

Assigns a permission scheme with a project. See [Managing project permissions](https://confluence.atlassian.com/x/yodKLg) for more information about permission schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg)

## Parameters
- `projectKeyOrId` [path] (required) string — The project ID or project key (case sensitive).
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Request Body
Content-Type: `application/json`
object:
  - `id` (required): integer(int64)

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

 *  the user does not have the necessary permission to edit the project's configuration.
 *  the Jira instance is Jira Core Free or Jira Software Free. Permission schemes cannot be assigned to projects on free plans.
- 404: Returned if the project or permission scheme is not found.
