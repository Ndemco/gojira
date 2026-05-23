# POST /rest/api/3/workflows/create/validation
**operationId:** `validateCreateWorkflows`
**Summary:** Validate create workflows

Validate the payload for bulk create workflows.

**[Permissions](#permissions) required:**

 *  *Administer Jira* project permission to create all, including global-scoped, workflows
 *  *Administer projects* project permissions to create project-scoped workflows

## Request Body
Content-Type: `application/json`
object:
  - `payload` (required): WorkflowCreateRequest
  - `validationOptions`: ValidationOptionsForCreate

## Responses
- 200: object:
  - `errors`: []WorkflowValidationError
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
