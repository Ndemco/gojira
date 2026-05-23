# GET /rest/api/3/project/{projectIdOrKey}/roledetails
**operationId:** `getProjectRoleDetails`
**Summary:** Get project role details

Returns all [project roles](https://support.atlassian.com/jira-cloud-administration/docs/manage-project-roles/) and the details for each role. Note that the list of project roles is common to all projects.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `currentMember` [query] boolean — Whether the roles should be filtered to include only those the user is assigned to.
- `excludeConnectAddons` [query] boolean — 
- `excludeOtherServiceRoles` [query] boolean — Do not return the default JSM company-managed space from CSM spaces, or the default CSM roles from JSM spaces.

## Responses
- 200: []object:
  - `admin`: boolean
  - `default`: boolean
  - `description`: string
  - `id`: integer(int64)
  - `name`: string
  - `roleConfigurable`: boolean
  - `scope`: allOf(Scope)
  - `self`: string(uri)
  - `translatedName`: string
  - `type`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or if the user does not have the necessary permissions for the project.
