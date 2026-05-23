# POST /rest/api/3/field/{fieldKey}/option
**operationId:** `createIssueFieldOption`
**Summary:** Create issue field option

Creates an option for a select list issue field.

Note that this operation **only works for issue field select list options added by Connect apps**, it cannot be used with issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource.

Each field can have a maximum of 10000 options, and each option can have a maximum of 10000 scopes.

**[Permissions](#permissions) required:** *Administer Jira* [global per

## Parameters
- `fieldKey` [path] (required) string — The field key is specified in the following format: **$(app-key)\_\_$(field-key)**. For example, *example-add-on\_\_exam

## Request Body
Content-Type: `application/json`
object:
  - `config`: IssueFieldOptionConfiguration
  - `properties`: object
  - `value` (required): string

## Responses
- 200: object:
  - `config`: IssueFieldOptionConfiguration
  - `id` (required): integer(int64)
  - `properties`: object
  - `value` (required): string
- 400: Returned if the option is invalid.
- 403: Returned if the request is not authenticated as a Jira administrator or the app that provided the field.
- 404: Returned if the field is not found or does not support options.
