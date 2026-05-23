# GET /rest/api/3/filter/{id}
**operationId:** `getFilter`
**Summary:** Get filter

Returns a filter.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None, however, the filter is only returned where it is:

 *  owned by the user.
 *  shared with a group that the user is a member of.
 *  shared with a private project that the user has *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for.
 *  shared with a public project.
 *  shared with the public.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter to return.
- `expand` [query] string — Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-
- `overrideSharePermissions` [query] boolean — EXPERIMENTAL: Whether share permissions are overridden to enable filters with any share permissions to be returned. Avai

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
- 400: Returned if the filter is not found or the user does not have permission to view it.
- 401: Returned if the authentication credentials are incorrect or missing.
