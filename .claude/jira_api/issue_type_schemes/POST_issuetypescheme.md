# POST /rest/api/3/issuetypescheme
**operationId:** `createIssueTypeScheme`
**Summary:** Create issue type scheme

Creates an issue type scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `defaultIssueTypeId`: string
  - `description`: string
  - `issueTypeIds` (required): []string
  - `name` (required): string

## Responses
- 201: object:
  - `issueTypeSchemeId` (required): string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 409: any
