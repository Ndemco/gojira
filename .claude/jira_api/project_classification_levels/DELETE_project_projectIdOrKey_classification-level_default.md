# DELETE /rest/api/3/project/{projectIdOrKey}/classification-level/default
**operationId:** `removeDefaultProjectClassification`
**Summary:** Remove the default data classification level from a project

Remove the default data classification level for a project.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case-sensitive).

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the user does not have the necessary permission.
- 404: Returned if the project is not found.
