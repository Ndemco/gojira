# GET /rest/api/3/screens/{screenId}/tabs
**operationId:** `getAllScreenTabs`
**Summary:** Get all screen tabs

Returns the list of tabs for a screen.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).
 *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) when the project key is specified, providing that the screen is associated with the project through a Screen Scheme and Issue Type Screen Scheme.

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `projectKey` [query] string — The key of the project.

## Responses
- 200: []object:
  - `id`: integer(int64)
  - `name` (required): string
- 400: Returned if the screen ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen is not found.
