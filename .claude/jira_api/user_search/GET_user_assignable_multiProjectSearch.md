# GET /rest/api/3/user/assignable/multiProjectSearch
**operationId:** `findBulkAssignableUsers`
**Summary:** Find users assignable to projects

Returns a list of users who can be assigned issues in one or more projects. The list may be restricted to users whose attributes match a string.

This operation takes the users in the range defined by `startAt` and `maxResults`, up to the thousandth user, and then returns only the users from that range that can be assigned issues in the projects. This means the operation usually returns fewer users than specified in `maxResults`. To get all the users who can be assigned issues in the projects, u

## Parameters
- `query` [query] string — A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] string — A query string that is matched exactly against user `accountId`. Required, unless `query` is specified.
- `projectKeys` [query] (required) string — A list of project keys (case sensitive). This parameter accepts a comma-separated list.
- `startAt` [query] integer(int32) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.

## Responses
- 200: []object:
  - `accountId`: string
  - `accountType`: string
  - `active`: boolean
  - `appType`: string
  - `applicationRoles`: allOf(SimpleListWrapperApplicationRole)
  - `avatarUrls`: allOf(AvatarUrlsBean)
  - `displayName`: string
  - `emailAddress`: string
  - `expand`: string
  - `groups`: allOf(SimpleListWrapperGroupName)
  - `guest`: boolean
  - `key`: string
  - `locale`: string
  - `name`: string
  - `self`: string(uri)
  - `timeZone`: string
- 400: Returned if:

 *  `projectKeys` is missing.
 *  `query` or `accountId` is missing.
 *  `query` and `accountId` are provided.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if one or more of the projects is not found.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
