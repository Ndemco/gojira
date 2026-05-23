# PUT /rest/api/3/fieldconfigurationscheme/{id}
**operationId:** `updateFieldConfigurationScheme`
**Summary:** Update field configuration scheme

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Updates a field configuration scheme.

This operation can only update field configuration schemes used in company-managed (classic) projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the field configuration scheme.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
