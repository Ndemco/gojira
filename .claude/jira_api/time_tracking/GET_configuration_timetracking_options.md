# GET /rest/api/3/configuration/timetracking/options
**operationId:** `getSharedTimeTrackingConfiguration`
**Summary:** Get time tracking settings

Returns the time tracking settings. This includes settings such as the time format, default time unit, and others. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: object:
  - `defaultUnit` (required): string
  - `timeFormat` (required): string
  - `workingDaysPerWeek` (required): number(double)
  - `workingHoursPerDay` (required): number(double)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
