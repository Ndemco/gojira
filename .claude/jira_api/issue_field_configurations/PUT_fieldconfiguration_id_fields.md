# PUT /rest/api/3/fieldconfiguration/{id}/fields
**operationId:** `updateFieldConfigurationItems`
**Summary:** Update field configuration items

Deprecated, use [ Field schemes](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-field-schemes/#api-group-field-schemes) which supports field association schemes.

Updates fields in a field configuration. The properties of the field configuration fields provided override the existing values.

This operation can only update field configurations used in company-managed (classic) projects.

The operation can set the renderer for text fields to the default text renderer (`text-

## Parameters
- `id` [path] (required) integer(int64) — The ID of the field configuration.

## Request Body
Content-Type: `application/json`
object:
  - `fieldConfigurationItems` (required): []FieldConfigurationItem

## Responses
- 204: any
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the field configuration is not found.
