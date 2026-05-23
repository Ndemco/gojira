# GET /rest/api/3/project/{projectIdOrKey}/components
**operationId:** `getProjectComponents`
**Summary:** Get project components

Returns all components in a project. See the [Get project components paginated](#api-rest-api-3-project-projectIdOrKey-component-get) resource if you want to get a full list of components with pagination.

If your project uses Compass components, this API will return a paginated list of Compass components that are linked to issues in that project.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https://confluence.atla

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `componentSource` [query] string — The source of the components to return. Can be `jira` (default), `compass` or `auto`. When `auto` is specified, the API 

## Responses
- 200: []object:
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
- 404: Returned if the project is not found or the user does not have permission to view it.
