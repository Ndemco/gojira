# GET /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties
**operationId:** `getDashboardItemPropertyKeys`
**Summary:** Get dashboard item property keys

Returns the keys of all properties for a dashboard item.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** The user must have read permission of the dashboard or have the dashboard shared with them.

## Parameters
- `dashboardId` [path] (required) string — The ID of the dashboard.
- `itemId` [path] (required) string — The ID of the dashboard item.

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
