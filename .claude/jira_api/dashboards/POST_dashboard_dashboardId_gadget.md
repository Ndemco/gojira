# POST /rest/api/3/dashboard/{dashboardId}/gadget
**operationId:** `addGadget`
**Summary:** Add gadget to dashboard

Adds a gadget to a dashboard.

**[Permissions](#permissions) required:** None.

## Parameters
- `dashboardId` [path] (required) integer(int64) — The ID of the dashboard.

## Request Body
Content-Type: `application/json`
object:
  - `color`: string
  - `ignoreUriAndModuleKeyValidation`: boolean
  - `moduleKey`: string
  - `position`: allOf(DashboardGadgetPosition)
  - `title`: string
  - `uri`: string

## Responses
- 200: object:
  - `color` (required): string
  - `id` (required): integer(int64)
  - `moduleKey`: string
  - `position` (required): allOf(DashboardGadgetPosition)
  - `title` (required): string
  - `uri`: string
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
