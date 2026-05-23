# GET /rest/api/3/license/approximateLicenseCount/product/{applicationKey}
**operationId:** `getApproximateApplicationLicenseCount`
**Summary:** Get approximate application license count

Returns the total approximate number of user accounts for a single Jira license. Note that this information is cached with a 7-day lifecycle and could be stale at the time of call.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `applicationKey` [path] (required) string — The ID of the application, represents a specific version of Jira.

## Responses
- 200: object:
  - `key`: string
  - `value`: string
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
