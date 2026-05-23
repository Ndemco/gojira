# POST /rest/api/3/notificationscheme
**operationId:** `createNotificationScheme`
**Summary:** Create notification scheme

Creates a notification scheme with notifications. You can create up to 1000 notifications per request.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string
  - `notificationSchemeEvents`: []NotificationSchemeEventDetails

## Responses
- 201: object:
  - `id` (required): string
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
