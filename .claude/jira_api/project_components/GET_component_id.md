# GET /rest/api/3/component/{id}
**operationId:** `getComponent`
**Summary:** Get component

Returns a component.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for project containing the component.

## Parameters
- `id` [path] (required) string — The ID of the component.

## Responses
- 200: object:
  - `ari`: string
  - `assignee`: allOf(User)
  - `assigneeType`: string
  - `description`: string
  - `id`: string
  - `isAssigneeTypeValid`: boolean
  - `lead`: allOf(User)
  - `leadAccountId`: string
  - `leadUserName`: string
  - `metadata`: object
  - `name`: string
  - `project`: string
  - `projectId`: integer(int64)
  - `realAssignee`: allOf(User)
  - `realAssigneeType`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the component is not found or the user does not have permission to browse the project containing the component.
