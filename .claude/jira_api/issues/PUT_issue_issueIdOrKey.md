# PUT /rest/api/3/issue/{issueIdOrKey}
**operationId:** `editIssue`
**Summary:** Edit issue

Edits an issue. Issue properties may be updated as part of the edit. Please note that issue transition is not supported and is ignored here. To transition an issue, please use [Transition issue](#api-rest-api-3-issue-issueIdOrKey-transitions-post).

The edits to the issue's fields are defined using `update` and `fields`. The fields that can be edited are determined using [ Get edit issue metadata](#api-rest-api-3-issue-issueIdOrKey-editmeta-get).

The parent field may be set by key or ID. For st

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `notifyUsers` [query] boolean — Whether a notification email about the issue update is sent to all watchers. To disable the notification, administer Jir
- `overrideScreenSecurity` [query] boolean — Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users wit
- `overrideEditableFlag` [query] boolean — Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users
- `returnIssue` [query] boolean — Whether the response should contain the issue with fields edited in this request. The returned issue will have the same 
- `expand` [query] string — The Get issue API expand parameter to use in the response if the `returnIssue` parameter is `true`.

## Request Body
Content-Type: `application/json`
object:
  - `fields`: object
  - `historyMetadata`: allOf(HistoryMetadata)
  - `properties`: []EntityProperty
  - `transition`: allOf(IssueTransition)
  - `update`: object

## Responses
- 200: any
- 204: any
- 400: Returned if:

 *  the request body is missing.
 *  the user does not have the necessary permission to edit one or more fields.
 *  the request includes one or more fields that are not found or are not associated with the issue's edit screen.
 *  the request includes an invalid transition.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user uses `overrideScreenSecurity` or `overrideEditableFlag` but doesn't have the necessary permission.
- 404: Returned if the issue is not found or the user does not have permission to view it.
- 409: Returned if the issue could not be updated due to a conflicting update.
- 422: Returned if a configuration problem prevents the issue being updated.
