# GET /rest/api/3/notificationscheme
**operationId:** `getNotificationSchemes`
**Summary:** Get notification schemes paginated

Returns a [paginated](#pagination) list of [notification schemes](https://confluence.atlassian.com/x/8YdKLg) ordered by the display name.

*Note that you should allow for events without recipients to appear in responses.*

**[Permissions](#permissions) required:** Permission to access Jira, however, the user must have permission to administer at least one project associated with a notification scheme for it to be returned.

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of notification schemes IDs to be filtered by
- `projectId` [query] []string — The list of projects IDs to be filtered by
- `onlyDefault` [query] boolean — When set to true, returns only the default notification scheme. If you provide project IDs not associated with the defau
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []NotificationScheme
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
