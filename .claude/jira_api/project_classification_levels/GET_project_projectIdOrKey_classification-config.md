# GET /rest/api/3/project/{projectIdOrKey}/classification-config
**operationId:** `getProjectClassificationConfig`
**Summary:** Get the classification configuration for a project

Returns the consolidated classification configuration for a project's admin settings page.

This includes permitted classification levels (with status), the project's default classification level, the organization's default classification level, and the container override setting.

**[Permissions](#permissions) required:**

 *  *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer projects* [project permission](https://confluence.atlas

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case-sensitive).

## Responses
- 200: any
- 401: Returned if the user does not have the necessary permission.
- 404: Returned if the project is not found or the feature is disabled.
