# GET /rest/api/3/filter/defaultShareScope
**operationId:** `getDefaultShareScope`
**Summary:** Get default share scope

Returns the default sharing settings for new filters and dashboards for a user.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: object:
  - `scope` (required): string
- 401: Returned if the authentication credentials are incorrect or missing.
