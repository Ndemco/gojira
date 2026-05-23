# GET /rest/api/3/issueLinkType
**operationId:** `getIssueLinkTypes`
**Summary:** Get issue link types

Returns a list of all issue link types.

To use this operation, the site must have [issue linking](https://confluence.atlassian.com/x/yoXKM) enabled.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for a project in the site.

## Responses
- 200: object:
  - `issueLinkTypes`: []IssueLinkType
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if issue linking is disabled.
