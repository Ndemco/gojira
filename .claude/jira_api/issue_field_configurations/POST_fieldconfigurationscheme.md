# POST /rest/api/3/fieldconfigurationscheme
**operationId:** `createFieldConfigurationScheme`
**Summary:** Create field configuration scheme

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Creates a field configuration scheme.

This operation can only create field configuration schemes used in company-managed (classic) projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string

## Responses
- 201: object:
  - `description`: string
  - `id` (required): string
  - `name` (required): string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
