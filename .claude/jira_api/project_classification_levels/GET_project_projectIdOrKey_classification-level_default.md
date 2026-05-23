# GET /rest/api/3/project/{projectIdOrKey}/classification-level/default
**operationId:** `getDefaultProjectClassification`
**Summary:** Get the default data classification level of a project

Returns the default data classification for a project.

**[Permissions](#permissions) required:**

 *  *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case-sensitive).

## Responses
- 200: any
- 401: Returned if the user does not have the necessary permission.
- 404: Returned if the project is not found.
