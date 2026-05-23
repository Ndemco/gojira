# GET /rest/api/3/groupuserpicker
**operationId:** `findUsersAndGroups`
**Summary:** Find users and groups

Returns a list of users and groups matching a string. The string is used:

 *  for users, to find a case-insensitive match with display name and e-mail address. Note that if a user has hidden their email address in their user profile, partial matches of the email address will not find the user. An exact match is required.
 *  for groups, to find a case-sensitive match with group name.

For example, if the string *tin* is used, records with the display name *Tina*, email address *sarah@tinplatetr

## Parameters
- `query` [query] (required) string — The search string.
- `maxResults` [query] integer(int32) — The maximum number of items to return in each list.
- `showAvatar` [query] boolean — Whether the user avatar should be returned. If an invalid value is provided, the default value is used.
- `fieldId` [query] string — The custom field ID of the field this request is for.
- `projectId` [query] []string — The ID of a project that returned users and groups must have permission to view. To include multiple projects, provide a
- `issueTypeId` [query] []string — The ID of an issue type that returned users and groups must have permission to view. To include multiple issue types, pr
- `avatarSize` [query] string — The size of the avatar to return. If an invalid value is provided, the default value is used.
- `caseInsensitive` [query] boolean — Whether the search for groups should be case insensitive.
- `excludeConnectAddons` [query] boolean — Whether Connect app users and groups should be excluded from the search results. If an invalid value is provided, the de

## Responses
- 200: object:
  - `groups`: FoundGroups
  - `users`: FoundUsers
- 400: Returned if the query parameter is not provided.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 429: Returned if the rate limit is exceeded. User search endpoints share a collective rate limit for the tenant, in addition to Jira's normal rate limiting you may receive a rate limit for user search. Please respect the Retry-After header.
