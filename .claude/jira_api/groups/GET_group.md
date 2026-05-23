# GET /rest/api/3/group
**operationId:** `getGroup`
**Summary:** Get group

This operation is deprecated, use [`group/member`](#api-rest-api-3-group-member-get).

Returns all users in a group.

**[Permissions](#permissions) required:** either of:

 *  *Browse users and groups* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended to identify a group.  
- `groupId` [query] string — The ID of the group. This parameter cannot be used with the `groupName` parameter.
- `expand` [query] string — List of fields to expand.

## Responses
- 200: object:
  - `expand`: string
  - `groupId`: string
  - `name`: string
  - `self`: string(uri)
  - `users`: allOf(PagedListUserDetailsApplicationUser)
- 400: Returned if the group name is not specified.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the calling user does not have the Administer Jira global permission.
- 404: Returned if the group is not found.
