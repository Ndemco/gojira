# DELETE /rest/api/3/group
**operationId:** `removeGroup`
**Summary:** Remove group

Deletes a group.

**[Permissions](#permissions) required:** Site administration (that is, member of the *site-admin* strategic [group](https://confluence.atlassian.com/x/24xjL)).

## Parameters
- `groupname` [query] string — 
- `groupId` [query] string — The ID of the group. This parameter cannot be used with the `groupname` parameter.
- `swapGroup` [query] string — As a group's name can change, use of `swapGroupId` is recommended to identify a group.  
- `swapGroupId` [query] string — The ID of the group to transfer restrictions to. Only comments and worklogs are transferred. If restrictions are not tra

## Responses
- 200: Returned if the request is successful.
- 400: Returned if the group name is not specified.
- 401: Returned if the authentication credentials are incorrect or missing from the request.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the group is not found.
