# GET /rest/api/3/user
**operationId:** `getUser`
**Summary:** Get user

Returns a user.

Privacy controls are applied to the response based on the user's preferences. This could mean, for example, that the user's email address is hidden. See the [Profile visibility overview](https://developer.atlassian.com/cloud/jira/platform/profile-visibility/) for more details.

**[Permissions](#permissions) required:** *Browse users and groups* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `accountId` [query] string — The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e0
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `key` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `expand` [query] string — Use [expand](#expansion) to include additional information about users in the response. This parameter accepts a comma-s

## Responses
- 200: object:
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
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the calling user does not have the *Browse users and groups* global permission.
- 404: Returned if the user is not found.
