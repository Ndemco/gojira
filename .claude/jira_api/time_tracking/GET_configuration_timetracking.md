# GET /rest/api/3/configuration/timetracking
**operationId:** `getSelectedTimeTrackingImplementation`
**Summary:** Get selected time tracking provider

Returns the time tracking provider that is currently selected. Note that if time tracking is disabled, then a successful but empty response is returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: object:
  - `key` (required): string
  - `name`: string
  - `url`: string
- 204: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
