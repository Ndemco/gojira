# POST /rest/api/3/issuetype
**operationId:** `createIssueType`
**Summary:** Create issue type

Creates an issue type.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `hierarchyLevel`: integer(int32)
  - `name` (required): string
  - `type`: string

## Responses
- 201: object:
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
- 400: Returned if the request is invalid because:

 *  no content is sent.
 *  the issue type name exceeds 60 characters.
 *  a subtask issue type is requested on an instance where subtasks are disabled.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 409: Returned if the issue type name is in use.
