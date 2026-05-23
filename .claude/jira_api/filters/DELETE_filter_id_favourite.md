# DELETE /rest/api/3/filter/{id}/favourite
**operationId:** `deleteFavouriteForFilter`
**Summary:** Remove filter as favorite

Removes a filter as a favorite for the user. Note that this operation only removes filters visible to the user from the user's favorites list. For example, if the user favorites a public filter that is subsequently made private (and is therefore no longer visible on their favorites list) they cannot remove it from their favorites list.

**[Permissions](#permissions) required:** Permission to access Jira.

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
 *  the user does not have permission to view the filter.
