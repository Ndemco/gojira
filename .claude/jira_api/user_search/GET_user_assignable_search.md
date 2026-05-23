# GET /rest/api/3/user/assignable/search
**operationId:** `findAssignableUsers`
**Summary:** Find users assignable to issues

Returns a list of users that can be assigned to an issue. Use this operation to find the list of users who can be assigned to:

 *  a new issue, by providing the `projectKeyOrId`.
 *  an updated issue, by providing the `issueKey` or `issueId`.
 *  to an issue during a transition (workflow action), by providing the `issueKey` or `issueId` and the transition id in `actionDescriptorId`. You can obtain the IDs of an issue's valid transitions using the `transitions` option in the `expand` parameter o

## Parameters
- `query` [query] string — A query string that is matched against user attributes, such as `displayName`, and `emailAddress`, to find relevant user
- `sessionId` [query] string — The sessionId of this request. SessionId is the same until the assignee is set.
- `username` [query] string — This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/
- `accountId` [query] string — A query string that is matched exactly against user `accountId`. Required, unless `query` is specified.
- `project` [query] string — The project ID or project key (case sensitive). Required, unless `issueKey` or `issueId` is specified.
- `issueKey` [query] string — The key of the issue. Required, unless `issueId` or `project` is specified.
- `issueId` [query] string — The ID of the issue. Required, unless `issueKey` or `project` is specified.
- `startAt` [query] integer(int32) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return. This operation may return less than the maximum number of items even if more are 
- `actionDescriptorId` [query] integer(int32) — The ID of the transition.
- `recommend` [query] boolean — 
- `accountType` [query] []string — 
- `appType` [query] []string — 

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

 *  None of `issueKey`, `issueId` or `project` is present.
 *  `issueId` parameter is not valid.
 *  `query` or `accountId` is missing.
 *  `query` and `accountId` are provided.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project, issue, or transition is not found.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
