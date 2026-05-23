# GET /rest/api/3/user/search
**operationId:** `findUsers`
**Summary:** Find users

Returns a list of active users that match the search string and property.

This operation first applies a filter to match the search string and property, and then takes the filtered users in the range defined by `startAt` and `maxResults`, up to the thousandth user. To get all the users who match the search string and property, use [Get all users](#api-rest-api-3-users-search-get) and filter the records in your code.

This operation can be accessed anonymously.

Privacy controls are applied to t

## Parameters
- `query` [query] string — A query string that is matched against user attributes ( `displayName`, and `emailAddress`) to find relevant users. The 
- `username` [query] string — 
- `accountId` [query] string — A query string that is matched exactly against a user `accountId`. Required, unless `query` or `property` is specified.
- `startAt` [query] integer(int32) — The index of the first item to return in a page of filtered results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `property` [query] string — A query string used to search properties. Property keys are specified by path, so property keys containing dot (.) or eq

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

 *  `accountId`, `query` or `property` is missing.
 *  `query` and `accountId` are provided.
 *  `property` parameter is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
