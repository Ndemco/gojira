# POST /rest/api/3/issue/{issueIdOrKey}/notify
**operationId:** `notify`
**Summary:** Send notification for issue

Creates an email notification for an issue and adds it to the mail queue.

**[Permissions](#permissions) required:**

 *  *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — ID or key of the issue that the notification is sent for.

## Request Body
Content-Type: `application/json`
object:
  - `htmlBody`: string
  - `restrict`: allOf(NotificationRecipientsRestrictions)
  - `subject`: string
  - `textBody`: string
  - `to`: allOf(NotificationRecipients)

## Responses
- 204: any
- 400: Returned if:

 *  the recipient is the same as the calling user.
 *  the recipient is invalid. For example, the recipient is set to the assignee, but the issue is unassigned.
 *  the issueIdOrKey is of an invalid/null issue.
 *  the request is invalid. For example, required fields are missing or have invalid values.
- 403: Returned if:

 *  outgoing emails are disabled.
 *  no SMTP server is configured.
- 404: Returned if the issue is not found.
