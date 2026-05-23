# GET /rest/api/3/issuetype/{id}/alternatives
**operationId:** `getAlternativeIssueTypes`
**Summary:** Get alternative issue types

Returns a list of issue types that can be used to replace the issue type. The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `id` [path] (required) string — The ID of the issue type.

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
- 404: Returned if:

 *  the issue type is not found.
 *  the user does not have the required permissions.
