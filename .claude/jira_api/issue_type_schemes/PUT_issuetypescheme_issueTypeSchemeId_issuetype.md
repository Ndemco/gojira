# PUT /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype
**operationId:** `addIssueTypesToIssueTypeScheme`
**Summary:** Add issue types to issue type scheme

Adds issue types to an issue type scheme.

The added issue types are appended to the issue types list.

If any of the issue types exist in the issue type scheme, the operation fails and no issue types are added.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeSchemeId` [path] (required) integer(int64) — The ID of the issue type scheme.

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
