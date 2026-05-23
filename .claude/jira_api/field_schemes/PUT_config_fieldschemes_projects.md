# PUT /rest/api/3/config/fieldschemes/projects
**operationId:** `associateProjectsToFieldAssociationSchemes`
**Summary:** Associate projects to field schemes

Associate projects to field association schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object

## Responses
- 200: object:
  - `results` (required): []FieldSchemeToProjectsPartialFailure
- 204: any
- 207: object:
  - `results` (required): []FieldSchemeToProjectsPartialFailure
- 400: object:
  - `results` (required): []FieldSchemeToProjectsPartialFailure
- 401: any
- 403: any
