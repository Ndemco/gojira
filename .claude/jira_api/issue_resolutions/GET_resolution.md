# GET /rest/api/3/resolution
**operationId:** `getResolutions`
**Summary:** Get resolutions

Returns a list of all issue resolution values.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: []object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
