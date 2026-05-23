# POST /rest/api/3/group
**operationId:** `createGroup`
**Summary:** Create group

Creates a group.

**[Permissions](#permissions) required:** Site administration (that is, member of the *site-admin* [group](https://confluence.atlassian.com/x/24xjL)).

## Request Body
Content-Type: `application/json`
object:
  - `name` (required): string

## Responses
- 201: object:
  - `expand`: string
  - `groupId`: string
  - `name`: string
  - `self`: string(uri)
  - `users`: allOf(PagedListUserDetailsApplicationUser)
- 400: Returned if group name is not specified or the group name is in use.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
