# GET /rest/api/3/field/{fieldId}/context/defaultValue
**operationId:** `getDefaultValues`
**Summary:** Get custom field contexts default values

Returns a [paginated](#pagination) list of defaults for a custom field. The results can be filtered by `contextId`, otherwise all values are returned. If no defaults are set for a context, nothing is returned.  
The returned object depends on type of the custom field:

 *  `CustomFieldContextDefaultValueDate` (type `datepicker`) for date fields.
 *  `CustomFieldContextDefaultValueDateTime` (type `datetimepicker`) for date-time fields.
 *  `CustomFieldContextDefaultValueSingleOption` (type `optio

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field, for example `customfield\_10000`.
- `contextId` [query] []integer(int64) — The IDs of the contexts.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []CustomFieldContextDefaultValue
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
