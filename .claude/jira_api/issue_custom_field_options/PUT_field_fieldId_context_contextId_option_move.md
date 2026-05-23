# PUT /rest/api/3/field/{fieldId}/context/{contextId}/option/move
**operationId:** `reorderCustomFieldOptions`
**Summary:** Reorder custom field options (context)

Changes the order of custom field options or cascading options in a context.

This operation works for custom field options created in Jira or the operations from this resource. **To work with issue field select list options created for Connect apps use the [Issue custom field options (apps)](#api-group-issue-custom-field-options--apps-) operations.**

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Request Body
Content-Type: `application/json`
object:
  - `after`: string
  - `customFieldOptionIds` (required): []string
  - `position`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
