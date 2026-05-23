# POST /rest/api/3/workflow/history/list
**operationId:** `listWorkflowHistory`
**Summary:** List workflow history entries

Returns a list of workflow history entries for a specified workflow id.

**Note:** Stored workflow data expires after 60 days. Additionally, no data from before the 30th of October 2025 is available.

**[Permissions](#permissions) required:**

 *  *Administer Jira* global permission to access all, including project-scoped, workflows
 *  At least one of the *Administer projects* and *View (read-only) workflow* project permissions to access project-scoped workflows

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Request Body
Content-Type: `application/json`
object:
  - `workflowId`: string

## Responses
- 200: object:
  - `entries`: []WorkflowHistoryItemDTO
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
