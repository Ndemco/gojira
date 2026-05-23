# GET /rest/api/3/myself
**operationId:** `getCurrentUser`
**Summary:** Get current user

Returns details for the current user.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information about user in the response. This parameter accepts a comma-se

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
