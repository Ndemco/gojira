# GET /rest/api/3/resolution/{id}
**operationId:** `getResolution`
**Summary:** Get resolution

Returns an issue resolution value.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `id` [path] (required) string — The ID of the issue resolution value.

## Responses
- 200: object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue resolution value is not found.
