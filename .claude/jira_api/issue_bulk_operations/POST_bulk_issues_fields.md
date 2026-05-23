# POST /rest/api/3/bulk/issues/fields
**operationId:** `submitBulkEdit`
**Summary:** Bulk edit issues

Use this API to submit a bulk edit request and simultaneously edit multiple issues. There are limits applied to the number of issues and fields that can be edited. A single request can accommodate a maximum of 1000 issues (including subtasks) and 200 fields.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).
 *  Browse [project permission](https://support.atlassian.com/jira-clou

## Request Body
Content-Type: `application/json`
object:
  - `editedFieldsInput` (required): allOf(JiraIssueFields)
  - `selectedActions` (required): []string
  - `selectedIssueIdsOrKeys` (required): []string
  - `sendBulkNotification`: boolean

## Responses
- 201: object:
  - `taskId`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
