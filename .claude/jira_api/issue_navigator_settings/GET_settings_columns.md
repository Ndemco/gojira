# GET /rest/api/3/settings/columns
**operationId:** `getIssueNavigatorDefaultColumns`
**Summary:** Get issue navigator default columns

Returns the default issue navigator columns.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: []object:
  - `label`: string
  - `value`: string
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
