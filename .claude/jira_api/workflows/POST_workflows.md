# POST /rest/api/3/workflows
**operationId:** `readWorkflows`
**Summary:** Bulk get workflows

Returns a list of workflows and related statuses by providing workflow names, workflow IDs, or project and issue types.

**[Permissions](#permissions) required:**

 *  *Administer Jira* global permission to access all, including project-scoped, workflows
 *  At least one of the *Administer projects* and *View (read-only) workflow* project permissions to access project-scoped workflows

## Request Body
Content-Type: `application/json`
object:
  - `projectAndIssueTypes`: []ProjectAndIssueTypePair
  - `workflowIds`: []string
  - `workflowNames`: []string

## Responses
- 200: object:
  - `statuses`: []JiraWorkflowStatus
  - `workflows`: []JiraWorkflow
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
