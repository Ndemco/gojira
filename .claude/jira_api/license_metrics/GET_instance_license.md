# GET /rest/api/3/instance/license
**operationId:** `getLicense`
**Summary:** Get license

Returns licensing information about the Jira instance.

**[Permissions](#permissions) required:** None.

## Responses
- 200: object:
  - `applications` (required): []LicensedApplication
- 401: Returned if the authentication credentials are incorrect or missing.
