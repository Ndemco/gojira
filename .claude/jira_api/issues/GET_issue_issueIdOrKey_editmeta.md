# GET /rest/api/3/issue/{issueIdOrKey}/editmeta
**operationId:** `getEditIssueMeta`
**Summary:** Get edit issue metadata

Returns the edit screen fields for an issue that are visible to and editable by the user. Use the information to populate the requests in [Edit issue](#api-rest-api-3-issue-issueIdOrKey-put).

This endpoint will check for these conditions:

1.  Field is available on a field screen - through screen, screen scheme, issue type screen scheme, and issue type scheme configuration. `overrideScreenSecurity=true` skips this condition.
2.  Field is visible in the [field configuration](https://support.atla

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `overrideScreenSecurity` [query] boolean — Whether hidden fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permission](
- `overrideEditableFlag` [query] boolean — Whether non-editable fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permis

## Responses
- 200: object:
  - `fields`: object
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user uses an override parameter but doesn't have permission to do so.
- 404: Returned if the issue is not found or the user does not have permission to view it.
