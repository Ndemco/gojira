# GET /rest/api/3/field/{fieldKey}/option
**operationId:** `getAllIssueFieldOptions`
**Summary:** Get all issue field options

Returns a [paginated](#pagination) list of all the options of a select list issue field. A select list issue field is a type of [issue field](https://developer.atlassian.com/cloud/jira/platform/modules/issue-field/) that enables a user to select a value from a list of options.

Note that this operation **only works for issue field select list options added by Connect apps**, it cannot be used with issue field select list options created in Jira or using operations from the [Issue custom field op

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `fieldKey` [path] (required) string — The field key is specified in the following format: **$(app-key)\_\_$(field-key)**. For example, *example-add-on\_\_exam

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []IssueFieldOption
- 400: Returned if the field is not found or does not support options.
- 403: Returned if the request is not authenticated as a Jira administrator or the app that provided the field.
