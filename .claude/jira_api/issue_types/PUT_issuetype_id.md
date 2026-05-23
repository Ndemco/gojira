# PUT /rest/api/3/issuetype/{id}
**operationId:** `updateIssueType`
**Summary:** Update issue type

Updates the issue type.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) string — The ID of the issue type.

## Request Body
Content-Type: `application/json`
object:
  - `avatarId`: integer(int64)
  - `description`: string
  - `name`: string

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
- 400: Returned if the request is invalid because:

 *  no content is sent.
 *  the issue type name exceeds 60 characters.
 *  the avatar is not associated with this issue type.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the issue type is not found.
- 409: Returned if the issue type name is in use.
