# GET /rest/api/3/projectCategory
**operationId:** `getAllProjectCategories`
**Summary:** Get all project categories

Returns all project categories.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: []object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
