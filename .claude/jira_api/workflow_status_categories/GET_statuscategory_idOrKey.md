# GET /rest/api/3/statuscategory/{idOrKey}
**operationId:** `getStatusCategory`
**Summary:** Get status category

Returns a status category. Status categories provided a mechanism for categorizing [statuses](#api-rest-api-3-status-idOrName-get).

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `idOrKey` [path] (required) string — The ID or key of the status category.

## Responses
- 200: object:
  - `colorName`: string
  - `id`: integer(int64)
  - `key`: string
  - `name`: string
  - `self`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the status category is not found.
