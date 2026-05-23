# GET /rest/api/3/user/permission/search
**operationId:** `findUsersWithAllPermissions`
**Summary:** Find users with permissions

Returns a list of users who fulfill these criteria:

 *  their user attributes match a search string.
 *  they have a set of permissions for a project or issue.

If no search string is provided, a list of all users with the permissions is returned.

This operation takes the users in the range defined by `startAt` and `maxResults`, up to the thousandth user, and then returns only the users from that range that match the search string and have permission for the project or issue. This means the op

## Parameters
- `query` [query] string — A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] string — A query string that is matched exactly against user `accountId`. Required, unless `query` is specified.
- `permissions` [query] (required) string — A comma separated list of permissions. Permissions can be specified as any:
- `issueKey` [query] string — The issue key for the issue.
- `projectKey` [query] string — The project key for the project (case sensitive).
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

 *  `issueKey` or `projectKey` is missing.
 *  `query` or `accountId` is missing.
 *  `query` and `accountId` are provided.
 *  `permissions` is empty or contains an invalid entry.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the issue or project is not found.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
