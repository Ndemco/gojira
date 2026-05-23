# PUT /rest/api/3/filter/{id}/favourite
**operationId:** `setFavouriteForFilter`
**Summary:** Add filter as favorite

Add a filter as a favorite for the user.

**[Permissions](#permissions) required:** Permission to access Jira, however, the user can only favorite:

 *  filters owned by the user.
 *  filters shared with a group that the user is a member of.
 *  filters shared with a private project that the user has *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for.
 *  filters shared with a public project.
 *  filters shared with the public.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.
- `expand` [query] string — Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-

## Responses
- 200: object:
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
- 400: Returned if:

 *  the filter is not found.
 *  the user does not have permission to favorite the filter.
