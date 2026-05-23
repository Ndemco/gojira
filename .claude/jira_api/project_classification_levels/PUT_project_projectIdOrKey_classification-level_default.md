# PUT /rest/api/3/project/{projectIdOrKey}/classification-level/default
**operationId:** `updateDefaultProjectClassification`
**Summary:** Update the default data classification level of a project

Updates the default data classification level for a project.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case-sensitive).

## Request Body
Content-Type: `application/json`
object:
  - `id` (required): string

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the user does not have the necessary permission.
- 404: Returned if the project is not found.
