# POST /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/remove
**operationId:** `removeMappingsFromIssueTypeScreenScheme`
**Summary:** Remove mappings from issue type screen scheme

Removes issue type to screen scheme mappings from an issue type screen scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeScreenSchemeId` [path] (required) string — The ID of the issue type screen scheme.

## Request Body
Content-Type: `application/json`
object:
  - `issueTypeIds` (required): []string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
