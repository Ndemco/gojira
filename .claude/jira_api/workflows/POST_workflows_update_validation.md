# POST /rest/api/3/workflows/update/validation
**operationId:** `validateUpdateWorkflows`
**Summary:** Validate update workflows

Validate the payload for bulk update workflows.

**[Permissions](#permissions) required:**

 *  *Administer Jira* project permission to create all, including global-scoped, workflows
 *  *Administer projects* project permissions to create project-scoped workflows

## Request Body
Content-Type: `application/json`
object:
  - `payload` (required): WorkflowUpdateRequest
  - `validationOptions`: ValidationOptionsForUpdate

## Responses
- 200: object:
  - `errors`: []WorkflowValidationError
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
