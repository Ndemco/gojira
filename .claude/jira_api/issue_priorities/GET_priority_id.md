# GET /rest/api/3/priority/{id}
**operationId:** `getPriority`
**Summary:** Get priority

Returns an issue priority.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `id` [path] (required) string — The ID of the issue priority.

## Responses
- 200: object:
  - `avatarId`: integer(int64)
  - `description`: string
  - `iconUrl`: string
  - `id`: string
  - `isDefault`: boolean
  - `name`: string
  - `schemes`: allOf(ExpandPrioritySchemePage)
  - `self`: string
  - `statusColor`: string
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if the issue priority isn't found.
