# GET /rest/api/3/screens/{screenId}/tabs/{tabId}/fields
**operationId:** `getAllScreenTabFields`
**Summary:** Get all screen tab fields

Returns all fields for a screen tab.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) when the project key is specified, providing that the screen is associated with the project through a Screen Scheme and Issue Type Screen Scheme.

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.
- `projectKey` [query] string — The key of the project.

## Responses
- 200: []object:
  - `id`: string
  - `name`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen or screen tab is not found.
