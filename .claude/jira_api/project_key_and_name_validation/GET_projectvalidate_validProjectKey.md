# GET /rest/api/3/projectvalidate/validProjectKey
**operationId:** `getValidProjectKey`
**Summary:** Get valid project key

Validates a project key and, if the key is invalid or in use, generates a valid random string for the project key.

**[Permissions](#permissions) required:** None.

## Parameters
- `key` [query] string — The project key.

## Responses
- 200: string
- 401: Returned if the authentication credentials are incorrect.
