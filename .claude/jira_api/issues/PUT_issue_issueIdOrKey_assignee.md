# PUT /rest/api/3/issue/{issueIdOrKey}/assignee
**operationId:** `assignIssue`
**Summary:** Assign issue

Assigns an issue to a user. Use this operation when the calling user does not have the *Edit Issues* permission but has the *Assign issue* permission for the project that the issue is in.

If `name` or `accountId` is set to:

 *  `"-1"`, the issue is assigned to the default assignee for the project.
 *  `null`, the issue is set to unassigned.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse Projects* and *Assign Issues* [ project permission](ht

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue to be assigned.

## Request Body
Content-Type: `application/json`
object:
  - `accountId`: string
  - `accountType`: string
  - `active`: boolean
  - `appType`: string
  - `applicationRoles`: allOf(SimpleListWrapperApplicationRole)
  - `avatarUrls`: allOf(AvatarUrlsBean)
  - `displayName`: string
  - `emailAddress`: string
  - `expand`: string
  - `groups`: allOf(SimpleListWrapperGroupName)
  - `guest`: boolean
  - `key`: string
  - `locale`: string
  - `name`: string
  - `self`: string(uri)
  - `timeZone`: string

## Responses
- 204: any
- 400: Returned if:

 *  the user is not found.
 *  `name`, `key`, or `accountId` is missing.
 *  more than one of `name`, `key`, and `accountId` are provided.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the issue is not found.
