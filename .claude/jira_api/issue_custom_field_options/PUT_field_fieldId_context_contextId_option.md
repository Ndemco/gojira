# PUT /rest/api/3/field/{fieldId}/context/{contextId}/option
**operationId:** `updateCustomFieldOption`
**Summary:** Update custom field options (context)

Updates the options of a custom field.

If any of the options are not found, no options are updated. Options where the values in the request match the current values aren't updated and aren't reported in the response.

Note that this operation **only works for issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource**, it cannot be used with issue field select list options created by Connect apps.

*

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Request Body
Content-Type: `application/json`
object:
  - `options`: []CustomFieldOptionUpdate

## Responses
- 200: object:
  - `options`: []CustomFieldOptionUpdate
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
