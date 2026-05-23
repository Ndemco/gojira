# DELETE /rest/api/3/field/{fieldId}/context/{contextId}/option/{optionId}/issue
**operationId:** `replaceCustomFieldOption`
**Summary:** Replace custom field options

Replaces the options of a custom field.

Note that this operation **only works for issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource**, it cannot be used with issue field select list options created by Connect or Forge apps.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `replaceWith` [query] integer(int64) — The ID of the option that will replace the currently selected option.
- `jql` [query] string — A JQL query that specifies the issues to be updated. For example, *project=10000*.
- `fieldId` [path] (required) string — The ID of the custom field.
- `optionId` [path] (required) integer(int64) — The ID of the option to be deselected.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Responses
- 303: object:
  - `description`: string
  - `elapsedRuntime` (required): integer(int64)
  - `finished`: integer(int64)
  - `id` (required): string
  - `lastUpdate` (required): integer(int64)
  - `message`: string
  - `progress` (required): integer(int64)
  - `result`: allOf(RemoveOptionFromIssuesResult)
  - `self` (required): string(uri)
  - `started`: integer(int64)
  - `status` (required): string
  - `submitted` (required): integer(int64)
  - `submittedBy` (required): integer(int64)
- 400: Returned if the request is not valid.
- 403: any
- 404: Returned if the field is not found or does not support options, or the options to be replaced are not found.
