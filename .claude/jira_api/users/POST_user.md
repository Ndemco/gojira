# POST /rest/api/3/user
**operationId:** `createUser`
**Summary:** Create user

Creates a user. This resource is retained for legacy compatibility. As soon as a more suitable alternative is available this resource will be deprecated.

**Note:** This API does not support Forge apps.

If the user exists and has access to Jira, the operation returns a 201 status. If the user exists but does not have access to Jira, the operation returns a 400 status.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). The

## Request Body
Content-Type: `application/json`
object:
  - `applicationKeys`: []string
  - `displayName`: string
  - `emailAddress` (required): string
  - `key`: string
  - `name`: string
  - `password`: string
  - `products` (required): []string
  - `self`: string

## Responses
- 201: object:
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
- 400: Returned if the request is invalid, the user already exists but does not have access to jira, or the number of licensed users is exceeded.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
