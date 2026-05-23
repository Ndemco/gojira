# DELETE /rest/api/3/field/{fieldId}/context/{contextId}/option/{optionId}
**operationId:** `deleteCustomFieldOption`
**Summary:** Delete custom field options (context)

Deletes a custom field option.

Options with cascading options cannot be deleted without deleting the cascading options first.

This operation works for custom field options created in Jira or the operations from this resource. **To work with issue field select list options created for Connect apps use the [Issue custom field options (apps)](#api-group-issue-custom-field-options--apps-) operations.**

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://conflue

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context from which an option should be deleted.
- `optionId` [path] (required) integer(int64) — The ID of the option to delete.

## Responses
- 204: Returned if the option is deleted.
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
