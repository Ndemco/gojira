# GET /rest/api/3/applicationrole/{key}
**operationId:** `getApplicationRole`
**Summary:** Get application role

Returns an application role.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `key` [path] (required) string — The key of the application role. Use the [Get all application roles](#api-rest-api-3-applicationrole-get) operation to g

## Responses
- 200: object:
  - `defaultGroups`: []string
  - `defaultGroupsDetails`: []GroupName
  - `defined`: boolean
  - `groupDetails`: []GroupName
  - `groups`: []string
  - `hasUnlimitedSeats`: boolean
  - `key`: string
  - `name`: string
  - `numberOfSeats`: integer(int32)
  - `platform`: boolean
  - `remainingSeats`: integer(int32)
  - `selectedByDefault`: boolean
  - `userCount`: integer(int32)
  - `userCountDescription`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user is not an administrator.
- 404: Returned if the role is not found.
