# GET /rest/api/3/project/type/accessible
**operationId:** `getAllAccessibleProjectTypes`
**Summary:** Get licensed project types

Returns all [project types](https://confluence.atlassian.com/x/Var1Nw) with a valid license.

## Responses
- 200: []object:
  - `color`: string
  - `descriptionI18nKey`: string
  - `formattedKey`: string
  - `icon`: string
  - `key`: string
