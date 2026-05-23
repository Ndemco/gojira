# GET /rest/api/3/project/type
**operationId:** `getAllProjectTypes`
**Summary:** Get all project types

Returns all [project types](https://confluence.atlassian.com/x/Var1Nw), whether or not the instance has a valid license for each type.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Responses
- 200: []object:
  - `color`: string
  - `descriptionI18nKey`: string
  - `formattedKey`: string
  - `icon`: string
  - `key`: string
- 401: Returned if the authentication credentials are incorrect.
