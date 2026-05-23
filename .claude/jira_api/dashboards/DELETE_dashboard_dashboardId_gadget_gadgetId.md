# DELETE /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId}
**operationId:** `removeGadget`
**Summary:** Remove gadget from dashboard

Removes a dashboard gadget from a dashboard.

When a gadget is removed from a dashboard, other gadgets in the same column are moved up to fill the emptied position.

**[Permissions](#permissions) required:** None.

## Parameters
- `dashboardId` [path] (required) integer(int64) — The ID of the dashboard.
- `gadgetId` [path] (required) integer(int64) — The ID of the gadget.

## Responses
- 204: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
