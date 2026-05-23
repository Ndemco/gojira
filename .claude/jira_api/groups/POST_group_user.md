# POST /rest/api/3/group/user
**operationId:** `addUserToGroup`
**Summary:** Add user to group

Adds a user to a group.

**[Permissions](#permissions) required:** Site administration (that is, member of the *site-admin* [group](https://confluence.atlassian.com/x/24xjL)).

## Parameters
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended to identify a group.  
- `groupId` [query] string — The ID of the group. This parameter cannot be used with the `groupName` parameter.

## Request Body
Content-Type: `application/json`
object:
  - `accountId`: string
  - `name`: string

## Responses
- 201: object:
  - `expand`: string
  - `groupId`: string
  - `name`: string
  - `self`: string(uri)
  - `users`: allOf(PagedListUserDetailsApplicationUser)
- 400: Returned if:

 *  `groupname` is not provided.
 *  `accountId` is missing.
- 401: Returned if the authentication credentials are incorrect or missing from the request.
- 403: Returned if the calling user does not have the necessary permission.
- 404: Returned if the group or user are not found.
- 429: Returned if rate limiting is being enforced.
