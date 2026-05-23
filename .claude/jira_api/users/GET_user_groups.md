# GET /rest/api/3/user/groups
**operationId:** `getUserGroups`
**Summary:** Get user groups

Returns the groups to which a user belongs.

**[Permissions](#permissions) required:** *Browse users and groups* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `accountId` [query] (required) string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `key` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/

## Responses
- 200: []object:
  - `groupId`: string
  - `name`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the calling user does not have the *Browse users and groups* global permission.
- 404: Returned if the user is not found.
