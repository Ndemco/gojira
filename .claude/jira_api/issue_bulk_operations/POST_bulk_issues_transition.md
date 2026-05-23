# POST /rest/api/3/bulk/issues/transition
**operationId:** `submitBulkTransition`
**Summary:** Bulk transition issue statuses

Use this API to submit a bulk issue status transition request. You can transition multiple issues, alongside with their valid transition Ids. You can transition up to 1,000 issues in a single operation.

**[Permissions](#permissions) required:**

 *  Global bulk change [permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-global-permissions/).
 *  Transition [issues permission](https://support.atlassian.com/jira-cloud-administration/docs/permissions-for-company-managed

## Request Body
Content-Type: `application/json`
object:
  - `bulkTransitionInputs` (required): []BulkTransitionSubmitInput
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
