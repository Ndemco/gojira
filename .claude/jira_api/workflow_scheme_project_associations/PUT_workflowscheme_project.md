# PUT /rest/api/3/workflowscheme/project
**operationId:** `assignSchemeToProject`
**Summary:** Assign workflow scheme to project

Assigns a workflow scheme to a project. This operation is performed only when there are no issues in the project.

Workflow schemes can only be assigned to classic projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `projectId` (required): string
  - `workflowSchemeId`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
