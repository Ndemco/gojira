# PUT /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId}
**operationId:** `updateGadget`
**Summary:** Update gadget on dashboard

Changes the title, position, and color of the gadget on a dashboard.

**[Permissions](#permissions) required:** None.

## Parameters
- `dashboardId` [path] (required) integer(int64) — The ID of the dashboard.
- `gadgetId` [path] (required) integer(int64) — The ID of the gadget.

## Request Body
Content-Type: `application/json`
object:
  - `color`: string
  - `position`: allOf(DashboardGadgetPosition)
  - `title`: string

## Responses
- 204: any
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
