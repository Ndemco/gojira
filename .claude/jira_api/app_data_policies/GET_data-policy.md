# GET /rest/api/3/data-policy
**operationId:** `getPolicy`
**Summary:** Get data policy for the workspace

Returns data policy for the workspace.

## Responses
- 200: object:
  - `anyContentBlocked`: boolean
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
