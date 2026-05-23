# POST /rest/api/3/component
**operationId:** `createComponent`
**Summary:** Create component

Creates a component. Use components to provide containers for issues within a project. Use components to provide containers for issues within a project.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project in which the component is created or *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

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
- 201: object:
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
 *  `name` is not provided.
 *  `name` is over 255 characters in length.
 *  `projectId` is not provided.
 *  `assigneeType` is an invalid value.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to manage the project containing the component or does not have permission to administer Jira.
- 404: Returned if the project is not found or the user does not have permission to browse the project containing the component.
