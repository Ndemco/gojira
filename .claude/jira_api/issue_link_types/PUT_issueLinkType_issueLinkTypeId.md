# PUT /rest/api/3/issueLinkType/{issueLinkTypeId}
**operationId:** `updateIssueLinkType`
**Summary:** Update issue link type

Updates an issue link type.

To use this operation, the site must have [issue linking](https://confluence.atlassian.com/x/yoXKM) enabled.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueLinkTypeId` [path] (required) string — The ID of the issue link type.

## Request Body
Content-Type: `application/json`
object:
  - `id`: string
  - `inward`: string
  - `name`: string
  - `outward`: string
  - `self`: string(uri)

## Responses
- 200: object:
  - `id`: string
  - `inward`: string
  - `name`: string
  - `outward`: string
  - `self`: string(uri)
- 400: Returned if the issue link type ID or the request body are invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  issue linking is disabled.
 *  the issue link type is not found.
 *  the user does not have the required permissions.
