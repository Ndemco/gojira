# DELETE /rest/api/3/screens/{screenId}/tabs/{tabId}
**operationId:** `deleteScreenTab`
**Summary:** Delete screen tab

Deletes a screen tab.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.

## Responses
- 204: Returned if the request is successful.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen or screen tab is not found.
