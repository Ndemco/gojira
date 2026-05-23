# GET /rest/api/3/project/{projectKeyOrId}/permissionscheme
**operationId:** `getAssignedPermissionScheme`
**Summary:** Get assigned permission scheme

Gets the [permission scheme](https://confluence.atlassian.com/x/yodKLg) associated with the project.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

## Parameters
- `projectKeyOrId` [path] (required) string — The project ID or project key (case sensitive).
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

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
- 403: Returned if the user does not have permission to view the project's configuration.
- 404: Returned if the project is not found or the user does not have permission to view the project.
