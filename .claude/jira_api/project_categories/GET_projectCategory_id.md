# GET /rest/api/3/projectCategory/{id}
**operationId:** `getProjectCategoryById`
**Summary:** Get project category by ID

Returns a project category.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the project category.

## Responses
- 200: object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project category is not found.
