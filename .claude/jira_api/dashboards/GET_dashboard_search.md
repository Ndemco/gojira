# GET /rest/api/3/dashboard/search
**operationId:** `getDashboardsPaginated`
**Summary:** Search for dashboards

Returns a [paginated](#pagination) list of dashboards. This operation is similar to [Get dashboards](#api-rest-api-3-dashboard-get) except that the results can be refined to include dashboards that have specific attributes. For example, dashboards with a particular name. When multiple attributes are specified only filters matching all attributes are returned.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** The following dashboards that match the query para

## Parameters
- `dashboardName` [query] string — String used to perform a case-insensitive partial match with `name`.
- `accountId` [query] string — User account ID used to return dashboards with the matching `owner.accountId`. This parameter cannot be used with the `o
- `owner` [query] string — This parameter is deprecated because of privacy changes. Use `accountId` instead. See the [migration guide](https://deve
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended. Group name used to return dashboards that are shared with
- `groupId` [query] string — Group ID used to return dashboards that are shared with a group that matches `sharePermissions.group.groupId`. This para
- `projectId` [query] integer(int64) — Project ID used to returns dashboards that are shared with a project that matches `sharePermissions.project.id`.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `status` [query] string — The status to filter by. It may be active, archived or deleted.
- `expand` [query] string — Use [expand](#expansion) to include additional information about dashboard in the response. This parameter accepts a com

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Dashboard
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
