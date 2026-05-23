# GET /rest/api/3/project/{projectKeyOrId}/issuesecuritylevelscheme
**operationId:** `getProjectIssueSecurityScheme`
**Summary:** Get project issue security scheme

Returns the [issue security scheme](https://confluence.atlassian.com/x/J4lKLg) associated with the project.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or the *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

## Parameters
- `projectKeyOrId` [path] (required) string — The project ID or project key (case sensitive).

## Responses
- 200: object:
  - `defaultSecurityLevelId`: integer(int64)
  - `description`: string
  - `id`: integer(int64)
  - `levels`: []SecurityLevel
  - `name`: string
  - `self`: string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the project is visible to the user but the user doesn't have administrative permissions.
- 404: Returned if the project is not found or the user does not have permission to view it.
