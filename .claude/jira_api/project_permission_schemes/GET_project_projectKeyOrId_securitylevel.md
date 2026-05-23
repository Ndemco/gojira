# GET /rest/api/3/project/{projectKeyOrId}/securitylevel
**operationId:** `getSecurityLevelsForProject`
**Summary:** Get project issue security levels

Returns all [issue security](https://confluence.atlassian.com/x/J4lKLg) levels for the project that the user has access to.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [global permission](https://confluence.atlassian.com/x/x4dKLg) for the project, however, issue security levels are only returned for authenticated user with *Set Issue Security* [global permission](https://confluence.atlassian.com/x/x4dKLg) for the project.

## Parameters
- `projectKeyOrId` [path] (required) string — The project ID or project key (case sensitive).

## Responses
- 200: object:
  - `levels` (required): []SecurityLevel
- 404: Returned if the project is not found or the user does not have permission to view it.
