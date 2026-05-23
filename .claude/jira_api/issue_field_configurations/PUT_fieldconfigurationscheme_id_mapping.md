# PUT /rest/api/3/fieldconfigurationscheme/{id}/mapping
**operationId:** `setFieldConfigurationSchemeMapping`
**Summary:** Assign issue types to field configurations

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Assigns issue types to field configurations on field configuration scheme.

This operation can only modify field configuration schemes used in company-managed (classic) projects.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the field configuration scheme.

## Request Body
Content-Type: `application/json`
object:
  - `mappings` (required): []FieldConfigurationToIssueTypeMapping

## Responses
- 204: any
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the field configuration scheme, the field configuration, or the issue type is not found.
