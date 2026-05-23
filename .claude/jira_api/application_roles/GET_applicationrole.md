# GET /rest/api/3/applicationrole
**operationId:** `getAllApplicationRoles`
**Summary:** Get all application roles

Returns all application roles. In Jira, application roles are managed using the [Application access configuration](https://confluence.atlassian.com/x/3YxjL) page.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: []object:
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
