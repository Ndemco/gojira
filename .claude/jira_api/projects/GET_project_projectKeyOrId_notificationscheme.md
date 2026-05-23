# GET /rest/api/3/project/{projectKeyOrId}/notificationscheme
**operationId:** `getNotificationSchemeForProject`
**Summary:** Get project notification scheme

Gets a [notification scheme](https://confluence.atlassian.com/x/8YdKLg) associated with the project.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

## Parameters
- `projectKeyOrId` [path] (required) string — The project ID or project key (case sensitive).
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
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user is not an administrator.
