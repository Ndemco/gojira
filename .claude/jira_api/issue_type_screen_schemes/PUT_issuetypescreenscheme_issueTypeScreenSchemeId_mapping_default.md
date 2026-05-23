# PUT /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/default
**operationId:** `updateDefaultScreenScheme`
**Summary:** Update issue type screen scheme default screen scheme

Updates the default screen scheme of an issue type screen scheme. The default screen scheme is used for all unmapped issue types.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeScreenSchemeId` [path] (required) string — The ID of the issue type screen scheme.

## Request Body
Content-Type: `application/json`
object:
  - `screenSchemeId` (required): string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
