# GET /rest/api/3/project/type/{projectTypeKey}
**operationId:** `getProjectTypeByKey`
**Summary:** Get project type by key

Returns a [project type](https://confluence.atlassian.com/x/Var1Nw).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `projectTypeKey` [path] (required) string — The key of the project type.

## Responses
- 200: object:
  - `color`: string
  - `descriptionI18nKey`: string
  - `formattedKey`: string
  - `icon`: string
  - `key`: string
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if the project type is not found.
