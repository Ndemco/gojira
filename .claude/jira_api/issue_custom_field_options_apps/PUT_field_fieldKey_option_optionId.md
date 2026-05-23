# PUT /rest/api/3/field/{fieldKey}/option/{optionId}
**operationId:** `updateIssueFieldOption`
**Summary:** Update issue field option

Updates or creates an option for a select list issue field. This operation requires that the option ID is provided when creating an option, therefore, the option ID needs to be specified as a path and body parameter. The option ID provided in the path and body must be identical.

Note that this operation **only works for issue field select list options added by Connect apps**, it cannot be used with issue field select list options created in Jira or using operations from the [Issue custom field 

## Parameters
- `fieldKey` [path] (required) string — The field key is specified in the following format: **$(app-key)\_\_$(field-key)**. For example, *example-add-on\_\_exam
- `optionId` [path] (required) integer(int64) — The ID of the option to be updated.

## Request Body
Content-Type: `application/json`
object:
  - `config`: IssueFieldOptionConfiguration
  - `id` (required): integer(int64)
  - `properties`: object
  - `value` (required): string

## Responses
- 200: object:
  - `config`: IssueFieldOptionConfiguration
  - `id` (required): integer(int64)
  - `properties`: object
  - `value` (required): string
- 400: Returned if the option is invalid, or the *ID* in the request object does not match the *optionId* parameter.
- 403: Returned if the request is not authenticated as a Jira administrator or the app that provided the field.
- 404: Returned if field is not found.
