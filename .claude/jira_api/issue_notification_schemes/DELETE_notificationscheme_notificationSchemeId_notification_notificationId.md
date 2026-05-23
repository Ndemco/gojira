# DELETE /rest/api/3/notificationscheme/{notificationSchemeId}/notification/{notificationId}
**operationId:** `removeNotificationFromNotificationScheme`
**Summary:** Remove notification from notification scheme

Removes a notification from a notification scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `notificationSchemeId` [path] (required) string — The ID of the notification scheme.
- `notificationId` [path] (required) string — The ID of the notification.

## Responses
- 204: any
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
