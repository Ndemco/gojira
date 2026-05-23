# GET /rest/api/3/license/approximateLicenseCount
**operationId:** `getApproximateLicenseCount`
**Summary:** Get approximate license count

Returns the approximate number of user accounts across all Jira licenses. Note that this information is cached with a 7-day lifecycle and could be stale at the time of call.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Responses
- 200: object:
  - `key`: string
  - `value`: string
- 401: object
- 403: object
