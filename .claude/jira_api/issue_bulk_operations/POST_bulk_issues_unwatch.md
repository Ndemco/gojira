# POST /rest/api/3/bulk/issues/unwatch
**operationId:** `submitBulkUnwatch`
**Summary:** Bulk unwatch issues

Use this API to submit a bulk unwatch request. You can unwatch up to 1,000 issues in a single operation.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).
 *  Browse [project permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-project-permissions/) in all projects that contain the selected issues.
 *  If [issue-level security](https://confluence.atla

## Request Body
Content-Type: `application/json`
object:
  - `selectedIssueIdsOrKeys` (required): []string

## Responses
- 201: object:
  - `taskId`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
- 403: object:
  - `errors`: []ErrorMessage
