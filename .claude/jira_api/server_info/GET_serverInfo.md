# GET /rest/api/3/serverInfo
**operationId:** `getServerInfo`
**Summary:** Get Jira instance info

Returns information about the Jira instance.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Responses
- 200: object:
  - `baseUrl`: string
  - `buildDate`: string(date-time)
  - `buildNumber`: integer(int32)
  - `deploymentType`: string
  - `displayUrl`: string
  - `displayUrlConfluence`: string
  - `displayUrlServicedeskHelpCenter`: string
  - `healthChecks`: []HealthCheckResult
  - `scmInfo`: string
  - `serverTime`: string(date-time)
  - `serverTimeZone`: string
  - `serverTitle`: string
  - `version`: string
  - `versionNumbers`: []integer(int32)
- 401: Returned if the authentication credentials are incorrect.
