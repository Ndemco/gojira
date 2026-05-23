# PUT /rest/api/3/fieldconfiguration/{id}
**operationId:** `updateFieldConfiguration`
**Summary:** Update field configuration

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Updates a field configuration. The name and the description provided in the request override the existing values.

This operation can only update configurations used in company-managed (classic) projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian

## Parameters
- `id` [path] (required) integer(int64) — The ID of the field configuration.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string

## Responses
- 204: any
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the field configuration is not found.
