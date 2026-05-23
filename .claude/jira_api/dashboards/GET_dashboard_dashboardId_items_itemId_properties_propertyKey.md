# GET /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}
**operationId:** `getDashboardItemProperty`
**Summary:** Get dashboard item property

Returns the key and value of a dashboard item property.

A dashboard item enables an app to add user-specific information to a user dashboard. Dashboard items are exposed to users as gadgets that users can add to their dashboards. For more information on how users do this, see [Adding and customizing gadgets](https://confluence.atlassian.com/x/7AeiLQ).

When an app creates a dashboard item it registers a callback to receive the dashboard item ID. The callback fires whenever the item is rendered 

## Parameters
- `dashboardId` [path] (required) string — The ID of the dashboard.
- `itemId` [path] (required) string — The ID of the dashboard item.
- `propertyKey` [path] (required) string — The key of the dashboard item property.

## Responses
- 200: object:
  - `key`: string
  - `value`: any
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
