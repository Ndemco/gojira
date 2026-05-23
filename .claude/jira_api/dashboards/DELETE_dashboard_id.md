# DELETE /rest/api/3/dashboard/{id}
**operationId:** `deleteDashboard`
**Summary:** Delete dashboard

Deletes a dashboard.

**[Permissions](#permissions) required:** None

The dashboard to be deleted must be owned by the user.

## Parameters
- `id` [path] (required) string — The ID of the dashboard.

## Responses
- 204: Returned if the dashboard is deleted.
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
