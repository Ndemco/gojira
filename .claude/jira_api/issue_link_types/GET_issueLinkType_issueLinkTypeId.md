# GET /rest/api/3/issueLinkType/{issueLinkTypeId}
**operationId:** `getIssueLinkType`
**Summary:** Get issue link type

Returns an issue link type.

To use this operation, the site must have [issue linking](https://confluence.atlassian.com/x/yoXKM) enabled.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for a project in the site.

## Parameters
- `issueLinkTypeId` [path] (required) string — The ID of the issue link type.

## Responses
- 200: object:
  - `id`: string
  - `inward`: string
  - `name`: string
  - `outward`: string
  - `self`: string(uri)
- 400: Returned if the issue link type ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  issue linking is disabled.
 *  the issue link type is not found.
 *  the user does not have the required permissions.
