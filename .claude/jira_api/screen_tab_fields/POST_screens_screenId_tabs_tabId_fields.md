# POST /rest/api/3/screens/{screenId}/tabs/{tabId}/fields
**operationId:** `addScreenTabField`
**Summary:** Add screen tab field

Adds a field to a screen tab.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.
- `skipFieldAssociation` [query] boolean — 

## Request Body
Content-Type: `application/json`
object:
  - `fieldId` (required): string

## Responses
- 200: object:
  - `id`: string
  - `name`: string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen, screen tab, or field is not found.
