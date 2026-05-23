# PUT /rest/api/3/component/{id}
**operationId:** `updateComponent`
**Summary:** Update component

Updates a component. Any fields included in the request are overwritten. If `leadAccountId` is an empty string ("") the component lead is removed.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the component or *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) string — The ID of the component.

## Request Body
Content-Type: `application/json`
object:
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
- 400: Returned if:

 *  the user is not found.
 *  `assigneeType` is an invalid value.
 *  `name` is over 255 characters in length.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to manage the project containing the component or does not have permission to administer Jira.
- 404: Returned if the component is not found or the user does not have permission to browse the project containing the component.
