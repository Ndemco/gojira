# DELETE /rest/api/3/workflow/{entityId}
**operationId:** `deleteInactiveWorkflow`
**Summary:** Delete inactive workflow

Deletes a workflow.

The workflow cannot be deleted if it is:

 *  an active workflow.
 *  a system workflow.
 *  associated with any workflow scheme.
 *  associated with any draft workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `entityId` [path] (required) string — The entity ID of the workflow.

## Responses
- 204: Returned if the workflow is deleted.
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
