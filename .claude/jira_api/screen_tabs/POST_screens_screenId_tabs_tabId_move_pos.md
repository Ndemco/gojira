# POST /rest/api/3/screens/{screenId}/tabs/{tabId}/move/{pos}
**operationId:** `moveScreenTab`
**Summary:** Move screen tab

Moves a screen tab.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.
- `pos` [path] (required) integer(int32) — The position of tab. The base index is 0.

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen or screen tab is not found or the position is invalid.
