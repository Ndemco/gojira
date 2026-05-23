# GET /rest/api/3/dashboard/gadgets
**operationId:** `getAllAvailableDashboardGadgets`
**Summary:** Get available gadgets

Gets a list of all available gadgets that can be added to all dashboards.

**[Permissions](#permissions) required:** None.

## Responses
- 200: object:
  - `gadgets` (required): []AvailableDashboardGadget
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
