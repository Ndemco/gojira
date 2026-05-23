# GET /rest/api/3/group/member
**operationId:** `getUsersFromGroup`
**Summary:** Get users from group

Returns a [paginated](#pagination) list of all users in a group.

Note that users are ordered by username, however the username is not returned in the results due to privacy reasons.

**[Permissions](#permissions) required:** either of:

 *  *Browse users and groups* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended to identify a group.  
- `groupId` [query] string — The ID of the group. This parameter cannot be used with the `groupName` parameter.
- `includeInactiveUsers` [query] boolean — Include inactive users.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page (number should be between 1 and 50).

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []UserDetails
- 400: Returned if the group name is not specified.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the calling user does not have the Administer Jira global permission.
- 404: Returned if the group is not found.
