# GET /rest/api/3/user/viewissue/search
**operationId:** `findUsersWithBrowsePermission`
**Summary:** Find users with browse permission

Returns a list of users who fulfill these criteria:

 *  their user attributes match a search string.
 *  they have permission to browse issues.

Use this resource to find users who can browse:

 *  an issue, by providing the `issueKey`.
 *  any issue in a project, by providing the `projectKey`.

This operation takes the users in the range defined by `startAt` and `maxResults`, up to the thousandth user, and then returns only the users from that range that match the search string and have permis

## Parameters
- `query` [query] string — A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] string — A query string that is matched exactly against user `accountId`. Required, unless `query` is specified.
- `issueKey` [query] string — The issue key for the issue. Required, unless `projectKey` is specified.
- `projectKey` [query] string — The project key for the project (case sensitive). Required, unless `issueKey` is specified.
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
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue or project is not found.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
