# GET /rest/api/3/filter/my
**operationId:** `getMyFilters`
**Summary:** Get my filters

Returns the filters owned by the user. If `includeFavourites` is `true`, the user's visible favorite filters are also returned.

**[Permissions](#permissions) required:** Permission to access Jira, however, a favorite filters is only visible to the user where the filter is:

 *  owned by the user.
 *  shared with a group that the user is a member of.
 *  shared with a private project that the user has *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for.
 *  shar

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-
- `includeFavourites` [query] boolean — Include the user's favorite filters in the response.

## Responses
- 200: []object:
  - `approximateLastUsed`: string(date-time)
  - `description`: string
  - `editPermissions`: []SharePermission
  - `favourite`: boolean
  - `favouritedCount`: integer(int64)
  - `id`: string
  - `jql`: string
  - `name` (required): string
  - `owner`: allOf(User)
  - `searchUrl`: string(uri)
  - `self`: string(uri)
  - `sharePermissions`: []SharePermission
  - `sharedUsers`: allOf(UserList)
  - `subscriptions`: allOf(FilterSubscriptionsList)
  - `viewUrl`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
