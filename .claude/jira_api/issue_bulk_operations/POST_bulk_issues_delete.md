# POST /rest/api/3/bulk/issues/delete
**operationId:** `submitBulkDelete`
**Summary:** Bulk delete issues

Use this API to submit a bulk delete request. You can delete up to 1,000 issues in a single operation.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).
 *  Delete [issues permission](https://support.atlassian.com/jira-cloud-administration/docs/permissions-for-company-managed-projects/#Delete-issues/) in all projects that contain the selected issues.
 *  Browse [project permiss

## Request Body
Content-Type: `application/json`
object:
  - `selectedIssueIdsOrKeys` (required): []string
  - `sendBulkNotification`: boolean

## Responses
- 201: object:
  - `taskId`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
- 403: object:
  - `errors`: []ErrorMessage
