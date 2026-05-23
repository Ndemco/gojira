# DELETE /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}
**operationId:** `deleteDashboardItemProperty`
**Summary:** Delete dashboard item property

Deletes a dashboard item property.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** The user must have edit permission of the dashboard.

## Parameters
- `dashboardId` [path] (required) string — The ID of the dashboard.
- `itemId` [path] (required) string — The ID of the dashboard item.
- `propertyKey` [path] (required) string — The key of the dashboard item property.

## Responses
- 204: any
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
