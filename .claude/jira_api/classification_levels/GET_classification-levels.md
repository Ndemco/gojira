# GET /rest/api/3/classification-levels
**operationId:** `getAllUserDataClassificationLevels`
**Summary:** Get all classification levels

Returns all classification levels.

**[Permissions](#permissions) required:** None.

## Parameters
- `status` [query] []string — Optional set of statuses to filter by.
- `orderBy` [query] string — Ordering of the results by a given field. If not provided, values will not be sorted.

## Responses
- 200: object:
  - `classifications`: []DataClassificationTagBean
- 401: Returned if the authentication credentials are incorrect or missing.
