# GET /rest/api/3/issue/{issueIdOrKey}/remotelink
**operationId:** `getRemoteIssueLinks`
**Summary:** Get remote issue links

Returns the remote issue links for an issue. When a remote issue link global ID is provided the record with that global ID is returned, otherwise all remote issue links are returned. Where a global ID includes reserved URL characters these must be escaped in the request. For example, pass `system=http://www.mycompany.com/support&id=1` as `system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1`.

This operation requires [issue linking to be active](https://confluence.atlassian.com/x/yoXKM).



## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `globalId` [query] string — The global ID of the remote issue link.

## Responses
- 200: oneOf([]object:
  - `application`: allOf(Application)
  - `globalId`: string
  - `id`: integer(int64)
  - `object`: allOf(RemoteObject)
  - `relationship`: string
  - `self`: string(uri), object:
  - `application`: allOf(Application)
  - `globalId`: string
  - `id`: integer(int64)
  - `object`: allOf(RemoteObject)
  - `relationship`: string
  - `self`: string(uri))
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if issue linking is disabled.
- 404: Returned if the issue or remote issue link is not found or the user does not have permission to view the issue.
- 413: Returned if the per-issue limit for remote links has been breached.
