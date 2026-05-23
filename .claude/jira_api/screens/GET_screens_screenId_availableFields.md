# GET /rest/api/3/screens/{screenId}/availableFields
**operationId:** `getAvailableScreenFields`
**Summary:** Get available screen fields

Returns the fields that can be added to a tab on a screen.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.

## Responses
- 200: []object:
  - `id`: string
  - `name`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen is not found.
