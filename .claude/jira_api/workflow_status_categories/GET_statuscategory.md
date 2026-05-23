# GET /rest/api/3/statuscategory
**operationId:** `getStatusCategories`
**Summary:** Get all status categories

Returns a list of all status categories.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: []object:
  - `colorName`: string
  - `id`: integer(int64)
  - `key`: string
  - `name`: string
  - `self`: string
- 401: Returned if the authentication credentials are incorrect or missing.
