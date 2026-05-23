# GET /rest/api/3/field/{fieldKey}/option/suggestions/edit
**operationId:** `getSelectableIssueFieldOptions`
**Summary:** Get selectable issue field options

Returns a [paginated](#pagination) list of options for a select list issue field that can be viewed and selected by the user.

Note that this operation **only works for issue field select list options added by Connect apps**, it cannot be used with issue field select list options created in Jira or using operations from the [Issue custom field options](#api-group-Issue-custom-field-options) resource.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `projectId` [query] integer(int64) — Filters the results to options that are only available in the specified project.
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
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the field is not found or does not support options.
