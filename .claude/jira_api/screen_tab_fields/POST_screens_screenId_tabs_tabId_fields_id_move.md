# POST /rest/api/3/screens/{screenId}/tabs/{tabId}/fields/{id}/move
**operationId:** `moveScreenTabField`
**Summary:** Move screen tab field

Moves a screen tab field.

If `after` and `position` are provided in the request, `position` is ignored.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.
- `id` [path] (required) string — The ID of the field.

## Request Body
Content-Type: `application/json`
object:
  - `after`: string(uri)
  - `position`: string

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen, screen tab, or field is not found or the field can't be moved to the requested position.
