# GET /rest/api/3/issue/{issueIdOrKey}/remotelink/{linkId}
**operationId:** `getRemoteIssueLinkById`
**Summary:** Get remote issue link by ID

Returns a remote issue link for an issue.

This operation requires [issue linking to be active](https://confluence.atlassian.com/x/yoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `linkId` [path] (required) string — The ID of the remote issue link.

## Responses
- 200: object:
  - `application`: allOf(Application)
  - `globalId`: string
  - `id`: integer(int64)
  - `object`: allOf(RemoteObject)
  - `relationship`: string
  - `self`: string(uri)
- 400: Returned if the link ID is invalid or the remote issue link does not belong to the issue.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if issue linking is disabled.
- 404: Returned if the issue or remote issue link is not found or the user does not have permission to view the issue.
