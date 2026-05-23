# POST /rest/api/3/issueLinkType
**operationId:** `createIssueLinkType`
**Summary:** Create issue link type

Creates an issue link type. Use this operation to create descriptions of the reasons why issues are linked. The issue link type consists of a name and descriptions for a link's inward and outward relationships.

To use this operation, the site must have [issue linking](https://confluence.atlassian.com/x/yoXKM) enabled.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `id`: string
  - `inward`: string
  - `name`: string
  - `outward`: string
  - `self`: string(uri)

## Responses
- 201: object:
  - `id`: string
  - `inward`: string
  - `name`: string
  - `outward`: string
  - `self`: string(uri)
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  issue linking is disabled.
 *  the issue link type name is in use.
 *  the user does not have the required permissions.
