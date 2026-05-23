# GET /rest/api/3/notificationscheme/{id}
**operationId:** `getNotificationScheme`
**Summary:** Get notification scheme

Returns a [notification scheme](https://confluence.atlassian.com/x/8YdKLg), including the list of events and the recipients who will receive notifications for those events.

**[Permissions](#permissions) required:** Permission to access Jira, however, the user must have permission to administer at least one project associated with the notification scheme.

## Parameters
- `id` [path] (required) integer(int64) — The ID of the notification scheme. Use [Get notification schemes paginated](#api-rest-api-3-notificationscheme-get) to g
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Responses
- 200: object:
  - `description`: string
  - `expand`: string
  - `id`: integer(int64)
  - `name`: string
  - `notificationSchemeEvents`: []NotificationSchemeEvent
  - `projects`: []integer(int64)
  - `scope`: allOf(Scope)
  - `self`: string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the notification scheme is not found or the user does not have permission to view it.
