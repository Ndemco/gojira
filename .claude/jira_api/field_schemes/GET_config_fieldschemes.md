# GET /rest/api/3/config/fieldschemes
**operationId:** `getFieldAssociationSchemes`
**Summary:** Get field schemes

REST endpoint for retrieving a paginated list of field association schemes with optional filtering.

This endpoint allows clients to fetch field association schemes with optional filtering by project IDs and text queries. The response includes scheme details with navigation links and filter metadata when applicable.

Filtering Behavior:

 *  When projectId or query parameters are provided, the response includes matchedFilters metadata showing which filters were applied.
 *  When no filters are a

## Parameters
- `projectId` [query] []integer(int64) — (optional) List of project IDs to filter schemes by. If not provided, schemes from all projects are returned.
- `query` [query] string — (optional) Text filter for scheme name or description matching (case-insensitive). If not provided, no text filtering is
- `startAt` [query] integer(int64) — Zero-based index of the first item to return (default: 0)
- `maxResults` [query] integer(int32) — Maximum number of items to return per page (default: 50, max: 100)

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []GetFieldAssociationSchemeResponse
- 400: object
- 401: any
- 403: any
- 404: any
