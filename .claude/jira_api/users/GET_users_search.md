# GET /rest/api/3/users/search
**operationId:** `getAllUsers`
**Summary:** Get all users

Returns a list of all users, including active users, inactive users and previously deleted users that have an Atlassian account.

Privacy controls are applied to the response based on the users' preferences. This could mean, for example, that the user's email address is hidden. See the [Profile visibility overview](https://developer.atlassian.com/cloud/jira/platform/profile-visibility/) for more details.

**[Permissions](#permissions) required:** *Browse users and groups* [global permission](htt

## Parameters
- `startAt` [query] integer(int32) — The index of the first item to return.
- `maxResults` [query] integer(int32) — The maximum number of items to return (limited to 1000).
- `expand` [query] string — 

## Responses
- 200: []object:
  - `accountId`: string
  - `accountType`: string
  - `active`: boolean
  - `appType`: string
  - `applicationRoles`: allOf(SimpleListWrapperApplicationRole)
  - `avatarUrls`: allOf(AvatarUrlsBean)
  - `displayName`: string
  - `emailAddress`: string
  - `expand`: string
  - `groups`: allOf(SimpleListWrapperGroupName)
  - `guest`: boolean
  - `key`: string
  - `locale`: string
  - `name`: string
  - `self`: string(uri)
  - `timeZone`: string
- 400: Returned if the request is invalid.
- 403: Returned if the user doesn't have the necessary permission.
- 409: Returned if the request takes longer than 10 seconds or is interrupted.
