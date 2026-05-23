# POST /rest/api/3/field/{fieldId}/context/{contextId}/option
**operationId:** `createCustomFieldOption`
**Summary:** Create custom field options (context)

Creates options and, where the custom select field is of the type Select List (cascading), cascading options for a custom select field. The options are added to a context of the field.

The maximum number of options that can be created per request is 1000 and each field can have a maximum of 10000 options.

This operation works for custom field options created in Jira or the operations from this resource. **To work with issue field select list options created for Connect apps use the [Issue cust

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Request Body
Content-Type: `application/json`
object:
  - `options`: []CustomFieldOptionCreate

## Responses
- 200: object:
  - `options`: []CustomFieldContextOption
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
