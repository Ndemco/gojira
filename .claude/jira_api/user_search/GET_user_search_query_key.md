# GET /rest/api/3/user/search/query/key
**operationId:** `findUserKeysByQuery`
**Summary:** Find user keys by query

Finds users with a structured query and returns a [paginated](#pagination) list of user keys.

This operation takes the users in the range defined by `startAt` and `maxResults`, up to the thousandth user, and then returns only the users from that range that match the structured query. This means the operation usually returns fewer users than specified in `maxResults`. To get all the users who match the structured query, use [Get all users](#api-rest-api-3-users-search-get) and filter the records

## Parameters
- `query` [query] (required) string — The search query.
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResult` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []UserKey
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 408: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
