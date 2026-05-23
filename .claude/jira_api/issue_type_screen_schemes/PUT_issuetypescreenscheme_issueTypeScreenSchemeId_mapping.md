# PUT /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping
**operationId:** `appendMappingsForIssueTypeScreenScheme`
**Summary:** Append mappings to issue type screen scheme

Appends issue type to screen scheme mappings to an issue type screen scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeScreenSchemeId` [path] (required) string — The ID of the issue type screen scheme.

## Request Body
Content-Type: `application/json`
object:
  - `issueTypeMappings` (required): []IssueTypeScreenSchemeMapping

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.
- 404: any
- 409: any
