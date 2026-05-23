# PUT /rest/api/3/filter/{id}
**operationId:** `updateFilter`
**Summary:** Update filter

Updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.

**[Permissions](#permissions) required:** Permission to access Jira, however the user must own the filter.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter to update.
- `expand` [query] string — Use [expand](#expansion) to include additional information about filter in the response. This parameter accepts a comma-
- `overrideSharePermissions` [query] boolean — EXPERIMENTAL: Whether share permissions are overridden to enable the addition of any share permissions to filters. Avail

## Request Body
Content-Type: `application/json`
object:
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
- 400: Returned if the request object is invalid. For example, the `name` is not unique or the project ID is not specified for a project role share permission.
- 401: Returned if the authentication credentials are incorrect or missing.
