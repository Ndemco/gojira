# GET /rest/api/3/issuetype/project
**operationId:** `getIssueTypesForProject`
**Summary:** Get issue types for project

Returns issue types for a project.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) in the relevant project or *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `projectId` [query] (required) integer(int64) — The ID of the project.
- `level` [query] integer(int32) — The level of the issue type to filter by. Use:

## Responses
- 200: []object:
  - `avatarId`: integer(int64)
  - `description`: string
  - `entityId`: string(uuid)
  - `hierarchyLevel`: integer(int32)
  - `iconUrl`: string
  - `id`: string
  - `name`: string
  - `scope`: allOf(Scope)
  - `self`: string
  - `subtask`: boolean
- 400: Returned if the request is not valid.
- 404: Returned if:

 *  the project is not found.
 *  the user does not have the necessary permission.
