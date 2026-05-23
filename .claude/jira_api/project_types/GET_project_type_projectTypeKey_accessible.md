# GET /rest/api/3/project/type/{projectTypeKey}/accessible
**operationId:** `getAccessibleProjectTypeByKey`
**Summary:** Get accessible project type by key

Returns a [project type](https://confluence.atlassian.com/x/Var1Nw) if it is accessible to the user.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `projectTypeKey` [path] (required) string — The key of the project type.

## Responses
- 200: object:
  - `color`: string
  - `descriptionI18nKey`: string
  - `formattedKey`: string
  - `icon`: string
  - `key`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project type is not accessible to the user.
