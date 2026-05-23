# GET /rest/api/3/user/bulk/migration
**operationId:** `bulkGetUsersMigration`
**Summary:** Get account IDs for users

Returns the account IDs for the users specified in the `key` or `username` parameters. Note that multiple `key` or `username` parameters can be specified.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `username` [query] []string — Username of a user. To specify multiple users, pass multiple copies of this parameter. For example, `username=fred&usern
- `key` [query] []string — Key of a user. To specify multiple users, pass multiple copies of this parameter. For example, `key=fred&key=barney`. Re

## Responses
- 200: []object:
  - `accountId`: string
  - `key`: string
  - `username`: string
- 400: Returned if `key` or `username`
- 401: Returned if the authentication credentials are incorrect or missing.
