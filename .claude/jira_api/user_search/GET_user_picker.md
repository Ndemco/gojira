# GET /rest/api/3/user/picker
**operationId:** `findUsersForPicker`
**Summary:** Find users for picker

Returns a list of users whose attributes match the query term. The returned object includes the `html` field where the matched query term is highlighted with the HTML strong tag. A list of account IDs can be provided to exclude users from the results.

This operation takes the users in the range defined by `maxResults`, up to the thousandth user, and then returns only the users from that range that match the query term. This means the operation usually returns fewer users than specified in `maxR

## Parameters
- `query` [query] (required) string — A query string that is matched against user attributes, such as `displayName`, and `emailAddress`, to find relevant user
- `maxResults` [query] integer(int32) — The maximum number of items to return. The total number of matched users is returned in `total`.
- `showAvatar` [query] boolean — Include the URI to the user's avatar.
- `exclude` [query] []string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `excludeAccountIds` [query] []string — A list of account IDs to exclude from the search results. This parameter accepts a comma-separated list. Multiple accoun
- `avatarSize` [query] string — 
- `excludeConnectUsers` [query] boolean — 

## Responses
- 200: object:
  - `header`: string
  - `total`: integer(int32)
  - `users`: []UserPickerUser
- 400: Returned if `exclude` and `excludeAccountIds` are provided.
- 401: Returned if the authentication credentials are incorrect or missing.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
