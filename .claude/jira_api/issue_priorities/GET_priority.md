# GET /rest/api/3/priority
**operationId:** `getPriorities`
**Summary:** Get priorities

Returns the list of all issue priorities.

**[Permissions](#permissions) required:** Permission to access Jira.

## Responses
- 200: []object:
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
