# GET /rest/api/3/configuration/timetracking/list
**operationId:** `getAvailableTimeTrackingImplementations`
**Summary:** Get all time tracking providers

Returns all time tracking providers. By default, Jira only has one time tracking provider: *JIRA provided time tracking*. However, you can install other time tracking providers via apps from the Atlassian Marketplace. For more information on time tracking providers, see the documentation for the [ Time Tracking Provider](https://developer.atlassian.com/cloud/jira/platform/modules/time-tracking-provider/) module.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](htt

## Responses
- 200: []object:
  - `key` (required): string
  - `name`: string
  - `url`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
