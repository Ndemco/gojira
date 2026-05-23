# PUT /rest/api/3/issues/archive/export
**operationId:** `exportArchivedIssues`
**Summary:** Export archived issue(s)

Enables admins to retrieve details of all archived issues. Upon a successful request, the admin who submitted it will receive an email with a link to download a CSV file with the issue details.

Note that this API only exports the values of system fields and archival-specific fields (`ArchivedBy` and `ArchivedDate`). Custom fields aren't supported.

**[Permissions](#permissions) required:** Jira admin or site admin: [global permission](https://confluence.atlassian.com/x/x4dKLg)

**License requir

## Request Body
Content-Type: `application/json`
object:
  - `archivedBy`: []string
  - `archivedDateRange`: DateRangeFilterRequest
  - `issueTypes`: []string
  - `projects`: []string
  - `reporters`: []string

## Responses
- 202: object:
  - `fileUrl`: string
  - `payload`: string
  - `progress`: integer(int64)
  - `status`: string
  - `submittedTime`: string(date-time)
  - `taskId`: string
- 400: any
- 401: any
- 403: any
- 412: any
