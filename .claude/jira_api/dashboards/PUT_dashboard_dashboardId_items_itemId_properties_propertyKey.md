# PUT /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}
**operationId:** `setDashboardItemProperty`
**Summary:** Set dashboard item property

Sets the value of a dashboard item property. Use this resource in apps to store custom data against a dashboard item.

A dashboard item enables an app to add user-specific information to a user dashboard. Dashboard items are exposed to users as gadgets that users can add to their dashboards. For more information on how users do this, see [Adding and customizing gadgets](https://confluence.atlassian.com/x/7AeiLQ).

When an app creates a dashboard item it registers a callback to receive the dashbo

## Parameters
- `dashboardId` [path] (required) string — The ID of the dashboard.
- `itemId` [path] (required) string — The ID of the dashboard item.
- `propertyKey` [path] (required) string — The key of the dashboard item property. The maximum length is 255 characters. For dashboard items with a spec URI and no

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
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
