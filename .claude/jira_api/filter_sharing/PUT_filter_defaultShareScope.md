# PUT /rest/api/3/filter/defaultShareScope
**operationId:** `setDefaultShareScope`
**Summary:** Set default share scope

Sets the default sharing for new filters and dashboards for a user.

**[Permissions](#permissions) required:** Permission to access Jira.

## Request Body
Content-Type: `application/json`
object:
  - `scope` (required): string

## Responses
- 200: object:
  - `scope` (required): string
- 400: Returned if an invalid scope is set.
- 401: Returned if the authentication credentials are incorrect or missing.
