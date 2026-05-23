# PUT /rest/api/3/issue/{issueIdOrKey}/remotelink/{linkId}
**operationId:** `updateRemoteIssueLink`
**Summary:** Update remote issue link by ID

Updates a remote issue link for an issue.

Note: Fields without values in the request are set to null.

This operation requires [issue linking to be active](https://confluence.atlassian.com/x/yoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Link issues* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lK

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `linkId` [path] (required) string — The ID of the remote issue link.

## Request Body
Content-Type: `application/json`
object:
  - `application`: allOf(Application)
  - `globalId`: string
  - `object` (required): allOf(RemoteObject)
  - `relationship`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to link issues.
- 404: Returned if the issue or remote issue link is not found or the user does not have permission to view the issue.
