# GET /rest/api/3/filter/search
**operationId:** `getFiltersPaginated`
**Summary:** Search for filters

Returns a [paginated](#pagination) list of filters. Use this operation to get:

 *  specific filters, by defining `id` only.
 *  filters that match all of the specified attributes. For example, all filters for a user with a particular word in their name. When multiple attributes are specified only filters matching all attributes are returned.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None, however, only the following filters that match the query para

## Parameters
- `filterName` [query] string — String used to perform a case-insensitive partial match with `name`.
- `accountId` [query] string — User account ID used to return filters with the matching `owner.accountId`. This parameter cannot be used with `owner`.
- `owner` [query] string — This parameter is deprecated because of privacy changes. Use `accountId` instead. See the [migration guide](https://deve
- `groupname` [query] string — As a group's name can change, use of `groupId` is recommended to identify a group. Group name used to returns filters th
- `groupId` [query] string — Group ID used to returns filters that are shared with a group that matches `sharePermissions.group.groupId`. This parame
- `projectId` [query] integer(int64) — Project ID used to returns filters that are shared with a project that matches `sharePermissions.project.id`.
- `id` [query] []integer(int64) — The list of filter IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. 
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `expand` [query] string — Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-
- `overrideSharePermissions` [query] boolean — EXPERIMENTAL: Whether share permissions are overridden to enable filters with any share permissions to be returned. Avai
- `isSubstringMatch` [query] boolean — When `true` this will perform a case-insensitive substring match for the provided `filterName`. When `false` the filter 

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []FilterDetails
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
