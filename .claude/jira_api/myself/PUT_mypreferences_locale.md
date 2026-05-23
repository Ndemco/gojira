# PUT /rest/api/3/mypreferences/locale
**operationId:** `setLocale`
**Summary:** Set locale

Deprecated, use [ Update a user profile](https://developer.atlassian.com/cloud/admin/user-management/rest/#api-users-account-id-manage-profile-patch) from the user management REST API instead.

Sets the locale of the user. The locale must be one supported by the instance of Jira.

**[Permissions](#permissions) required:** Permission to access Jira.

## Request Body
Content-Type: `application/json`
object:
  - `locale`: string

## Responses
- 204: any
- 400: Returned if request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
