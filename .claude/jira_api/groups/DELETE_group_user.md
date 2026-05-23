# DELETE /rest/api/3/group/user
**operationId:** `removeUserFromGroup`
**Summary:** Remove user from group

Removes a user from a group.

**[Permissions](#permissions) required:** Site administration (that is, member of the *site-admin* [group](https://confluence.atlassian.com/x/24xjL)).

## Parameters
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended to identify a group.  
- `groupId` [query] string — The ID of the group. This parameter cannot be used with the `groupName` parameter.
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] (required) string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0

## Responses
- 200: Returned if the request is successful.
- 400: Returned if:

 *  `groupName` is missing.
 *  `accountId` is missing.
- 401: Returned if the authentication credentials are incorrect or missing from the request.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the group or user are not found.
- 429: Returned if rate limiting is being enforced.
