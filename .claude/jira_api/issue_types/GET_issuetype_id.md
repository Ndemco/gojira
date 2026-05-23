# GET /rest/api/3/issuetype/{id}
**operationId:** `getIssueType`
**Summary:** Get issue type

Returns an issue type.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) in a project the issue type is associated with or *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) string — The ID of the issue type.

## Responses
- 200: object:
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
- 400: Returned if the issue type ID is invalid.
- 404: Returned if:

 *  the issue type is not found.
 *  the user does not have the required permissions.
