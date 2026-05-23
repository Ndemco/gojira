# PUT /rest/api/3/fieldconfigurationscheme/project
**operationId:** `assignFieldConfigurationSchemeToProject`
**Summary:** Assign field configuration scheme to project

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Assigns a field configuration scheme to a project. If the field configuration scheme ID is `null`, the operation assigns the default field configuration scheme.

Field configuration schemes can only be assigned to classic projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](h

## Request Body
Content-Type: `application/json`
object:
  - `fieldConfigurationSchemeId`: string
  - `projectId` (required): string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
