# PUT /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}
**operationId:** `updateIssueTypeScreenScheme`
**Summary:** Update issue type screen scheme

Updates an issue type screen scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeScreenSchemeId` [path] (required) string — The ID of the issue type screen scheme.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
