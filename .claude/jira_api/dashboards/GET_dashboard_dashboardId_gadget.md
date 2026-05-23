# GET /rest/api/3/dashboard/{dashboardId}/gadget
**operationId:** `getAllGadgets`
**Summary:** Get gadgets

Returns a list of dashboard gadgets on a dashboard.

This operation returns:

 *  Gadgets from a list of IDs, when `id` is set.
 *  Gadgets with a module key, when `moduleKey` is set.
 *  Gadgets from a list of URIs, when `uri` is set.
 *  All gadgets, when no other parameters are set.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `dashboardId` [path] (required) integer(int64) — The ID of the dashboard.
- `moduleKey` [query] []string — The list of gadgets module keys. To include multiple module keys, separate module keys with ampersand: `moduleKey=key:on
- `uri` [query] []string — The list of gadgets URIs. To include multiple URIs, separate URIs with ampersand: `uri=/rest/example/uri/1&uri=/rest/exa
- `gadgetId` [query] []integer(int64) — The list of gadgets IDs. To include multiple IDs, separate IDs with ampersand: `gadgetId=10000&gadgetId=10001`.

## Responses
- 200: object:
  - `gadgets` (required): []DashboardGadget
- 401: Returned if the authentication credentials are incorrect.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
