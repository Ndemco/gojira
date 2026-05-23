# GET /rest/api/3/projectvalidate/key
**operationId:** `validateProjectKey`
**Summary:** Validate project key

Validates a project key by confirming the key is a valid string and not in use.

**[Permissions](#permissions) required:** None.

## Parameters
- `key` [query] string — The project key.

## Responses
- 200: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect.
