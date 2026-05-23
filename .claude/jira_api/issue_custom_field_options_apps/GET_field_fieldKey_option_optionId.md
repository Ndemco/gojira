# GET /rest/api/3/field/{fieldKey}/option/{optionId}
**operationId:** `getIssueFieldOption`
**Summary:** Get issue field option

Returns an option from a select list issue field.

Note that this operation **only works for issue field select list options added by Connect apps**, it cannot be used with issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). Jira permissions are not required for the app p

## Parameters
- `fieldKey` [path] (required) string — The field key is specified in the following format: **$(app-key)\_\_$(field-key)**. For example, *example-add-on\_\_exam
- `optionId` [path] (required) integer(int64) — The ID of the option to be returned.

## Responses
- 200: object:
  - `config`: IssueFieldOptionConfiguration
  - `id` (required): integer(int64)
  - `properties`: object
  - `value` (required): string
- 400: Returned if the field is not found or does not support options.
- 403: Returned if the request is not authenticated as a Jira administrator or the app that provided the field.
- 404: Returned if the option is not found.
