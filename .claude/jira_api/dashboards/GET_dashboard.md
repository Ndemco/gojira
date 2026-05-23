# GET /rest/api/3/dashboard
**operationId:** `getAllDashboards`
**Summary:** Get all dashboards

Returns a list of dashboards owned by or shared with the user. The list may be filtered to include only favorite or owned dashboards.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `filter` [query] string — The filter applied to the list of dashboards. Valid values are:
- `startAt` [query] integer(int32) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `dashboards`: []Dashboard
  - `maxResults`: integer(int32)
  - `next`: string
  - `prev`: string
  - `startAt`: integer(int32)
  - `total`: integer(int32)
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
