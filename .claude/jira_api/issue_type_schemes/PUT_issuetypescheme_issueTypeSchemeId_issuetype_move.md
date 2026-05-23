# PUT /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype/move
**operationId:** `reorderIssueTypesInIssueTypeScheme`
**Summary:** Change order of issue types

Changes the order of issue types in an issue type scheme.

The request body parameters must meet the following requirements:

 *  all of the issue types must belong to the issue type scheme.
 *  either `after` or `position` must be provided.
 *  the issue type in `after` must not be in the issue type list.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `issueTypeSchemeId` [path] (required) integer(int64) — The ID of the issue type scheme.

## Request Body
Content-Type: `application/json`
object:
  - `after`: string
  - `issueTypeIds` (required): []string
  - `position`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
