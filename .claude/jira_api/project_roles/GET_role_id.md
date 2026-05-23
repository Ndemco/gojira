# GET /rest/api/3/role/{id}
**operationId:** `getProjectRoleById`
**Summary:** Get project role by ID

Gets the project role details and the default actors associated with the role. The list of default actors is sorted by display name.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.

## Responses
- 200: object:
  - `actors`: []RoleActor
  - `admin`: boolean
  - `currentUserRole`: boolean
  - `default`: boolean
  - `description`: string
  - `id`: integer(int64)
  - `name`: string
  - `roleConfigurable`: boolean
  - `scope`: allOf(Scope)
  - `self`: string(uri)
  - `translatedName`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have administrative permissions.
- 404: Returned if the project role is not found.
