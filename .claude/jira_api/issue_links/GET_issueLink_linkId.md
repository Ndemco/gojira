# GET /rest/api/3/issueLink/{linkId}
**operationId:** `getIssueLink`
**Summary:** Get issue link

Returns an issue link.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse project* [project permission](https://confluence.atlassian.com/x/yodKLg) for all the projects containing the linked issues.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, permission to view both of the issues.

## Parameters
- `linkId` [path] (required) string — The ID of the issue link.

## Responses
- 200: object:
  - `id`: string
  - `inwardIssue` (required): allOf(LinkedIssue)
  - `outwardIssue` (required): allOf(LinkedIssue)
  - `self`: string(uri)
  - `type` (required): allOf(IssueLinkType)
- 400: Returned if the issue link ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  issue linking is disabled.
 *  the issue link is not found.
 *  the user doesn't have the required permissions.
