# GET /rest/api/3/user/bulk
**operationId:** `bulkGetUsers`
**Summary:** Bulk get users

Returns a [paginated](#pagination) list of the users specified by one or more account IDs.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `username` [query] []string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http
- `key` [query] []string — This parameter is no longer available and will be removed from the documentation soon. See the [deprecation notice](http
- `accountId` [query] (required) []string — The account ID of a user. To specify multiple users, pass multiple `accountId` parameters. For example, `accountId=5b10a

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []User
- 400: Returned if `accountID` is missing.
- 401: Returned if the authentication credentials are incorrect or missing.
