# DELETE /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype/{issueTypeId}
**operationId:** `removeIssueTypeFromIssueTypeScheme`
**Summary:** Remove issue type from issue type scheme

Removes an issue type from an issue type scheme.

This operation cannot remove:

 *  any issue type used by issues.
 *  any issue types from the default issue type scheme.
 *  the last standard issue type from an issue type scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeSchemeId` [path] (required) integer(int64) — The ID of the issue type scheme.
- `issueTypeId` [path] (required) integer(int64) — The ID of the issue type.

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
