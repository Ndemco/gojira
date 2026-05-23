# PUT /rest/api/3/issuetypescreenscheme/project
**operationId:** `assignIssueTypeScreenSchemeToProject`
**Summary:** Assign issue type screen scheme to project

Assigns an issue type screen scheme to a project.

Issue type screen schemes can only be assigned to classic projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `issueTypeScreenSchemeId`: string
  - `projectId`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
