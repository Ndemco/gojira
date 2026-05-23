# DELETE /rest/api/3/field/{fieldKey}/option/{optionId}/issue
**operationId:** `replaceIssueFieldOption`
**Summary:** Replace issue field option

Deselects an issue-field select-list option from all issues where it is selected. A different option can be selected to replace the deselected option. The update can also be limited to a smaller set of issues by using a JQL query.

Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can override the screen security configuration using `overrideScreenSecurity` and `overrideEditableFlag`.

This is an [asynchronous operation](#async). Th

## Parameters
- `replaceWith` [query] integer(int64) — The ID of the option that will replace the currently selected option.
- `jql` [query] string — A JQL query that specifies the issues to be updated. For example, *project=10000*.
- `overrideScreenSecurity` [query] boolean — Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users wit
- `overrideEditableFlag` [query] boolean — Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users
- `fieldKey` [path] (required) string — The field key is specified in the following format: **$(app-key)\_\_$(field-key)**. For example, *example-add-on\_\_exam
- `optionId` [path] (required) integer(int64) — The ID of the option to be deselected.

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
