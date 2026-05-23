# PUT /rest/api/3/configuration/timetracking/options
**operationId:** `setSharedTimeTrackingConfiguration`
**Summary:** Set time tracking settings

Sets the time tracking settings.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `defaultUnit` (required): string
  - `timeFormat` (required): string
  - `workingDaysPerWeek` (required): number(double)
  - `workingHoursPerDay` (required): number(double)

## Responses
- 200: object:
  - `defaultUnit` (required): string
  - `timeFormat` (required): string
  - `workingDaysPerWeek` (required): number(double)
  - `workingHoursPerDay` (required): number(double)
- 400: Returned if the request object is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
