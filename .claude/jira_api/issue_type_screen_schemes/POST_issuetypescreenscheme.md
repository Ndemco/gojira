# POST /rest/api/3/issuetypescreenscheme
**operationId:** `createIssueTypeScreenScheme`
**Summary:** Create issue type screen scheme

Creates an issue type screen scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `issueTypeMappings` (required): []IssueTypeScreenSchemeMapping
  - `name` (required): string

## Responses
- 201: object:
  - `id` (required): string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
- 409: any
