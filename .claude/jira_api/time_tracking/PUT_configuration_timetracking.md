# PUT /rest/api/3/configuration/timetracking
**operationId:** `selectTimeTrackingImplementation`
**Summary:** Select time tracking provider

Selects a time tracking provider.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `key` (required): string
  - `name`: string
  - `url`: string

## Responses
- 204: any
- 400: Returned if the time tracking provider is not found.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
