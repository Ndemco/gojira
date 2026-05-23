# GET /rest/api/3/project/{projectIdOrKey}/role
**operationId:** `getProjectRoles`
**Summary:** Get project roles for project

Returns a list of [project roles](https://support.atlassian.com/jira-cloud-administration/docs/manage-project-roles/) for the project returning the name and self URL for each role.

Note that all project roles are shared with all projects in Jira Cloud. See [Get all project roles](#api-rest-api-3-role-get) for more information.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodK

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).

## Responses
- 200: object
- 401: Returned if the authentication credentials are incorrect or missing or if the user lacks administrative permissions for the project.
- 404: Returned if the project is not found or or if the user does not have administrative permissions for the project.
