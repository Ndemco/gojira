# POST /rest/api/3/issue/{issueIdOrKey}/remotelink
**operationId:** `createOrUpdateRemoteIssueLink`
**Summary:** Create or update remote issue link

Creates or updates a remote issue link for an issue.

If a `globalId` is provided and a remote issue link with that global ID is found it is updated. Any fields without values in the request are set to null. Otherwise, the remote issue link is created.

This operation requires [issue linking to be active](https://confluence.atlassian.com/x/yoXKM).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Link issues* [project permission](

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.

## Request Body
Content-Type: `application/json`
object:
  - `application`: allOf(Application)
  - `globalId`: string
  - `object` (required): allOf(RemoteObject)
  - `relationship`: string

## Responses
- 200: object:
  - `id`: integer(int64)
  - `self`: string
- 201: object:
  - `id`: integer(int64)
  - `self`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to link issues.
- 404: Returned if the issue is not found or the user does not have permission to view the issue.
