# GET /rest/api/3/customFieldOption/{id}
**operationId:** `getCustomFieldOption`
**Summary:** Get custom field option

Returns a custom field option. For example, an option in a select list.

Note that this operation **only works for issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource**, it cannot be used with issue field select list options created by Connect apps.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** The custom field option is returned as follows:

 *  if the 

## Parameters
- `id` [path] (required) string — The ID of the custom field option.

## Responses
- 200: object:
  - `self`: string(uri)
  - `value`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the custom field option is not found.
 *  the user does not have permission to view the custom field.
