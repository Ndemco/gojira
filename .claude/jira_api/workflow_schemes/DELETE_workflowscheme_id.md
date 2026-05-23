# DELETE /rest/api/3/workflowscheme/{id}
**operationId:** `deleteWorkflowScheme`
**Summary:** Delete workflow scheme

Deletes a workflow scheme. Note that a workflow scheme cannot be deleted if it is active (that is, being used by at least one project).

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL a

## Responses
- 204: any
- 400: Returned if the scheme is active.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme is not found.
