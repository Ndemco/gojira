# PUT /rest/api/3/issuetypescheme/{issueTypeSchemeId}
**operationId:** `updateIssueTypeScheme`
**Summary:** Update issue type scheme

Updates an issue type scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeSchemeId` [path] (required) integer(int64) — The ID of the issue type scheme.

## Request Body
Content-Type: `application/json`
object:
  - `defaultIssueTypeId`: string
  - `description`: string
  - `name`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
