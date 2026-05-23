# POST /rest/api/3/workflows/create
**operationId:** `createWorkflows`
**Summary:** Bulk create workflows

Create workflows and related statuses.

**[Permissions](#permissions) required:**

 *  *Administer Jira* project permission to create all, including global-scoped, workflows
 *  *Administer projects* project permissions to create project-scoped workflows

## Request Body
Content-Type: `application/json`
object:
  - `scope`: WorkflowScope
  - `statuses`: []WorkflowStatusUpdate
  - `workflows`: []WorkflowCreate

## Responses
- 200: object:
  - `statuses`: []JiraWorkflowStatus
  - `workflows`: []JiraWorkflow
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 409: Returned if another workflow configuration update task is ongoing.
