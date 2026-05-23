# DELETE /rest/api/3/workflowscheme/{id}/draft
**operationId:** `deleteWorkflowSchemeDraft`
**Summary:** Delete draft workflow scheme

Deletes a draft workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the active workflow scheme that the draft was created from.

## Responses
- 204: Returned if the request is successful.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission..
- 404: Returned if:

 *  the original active workflow scheme is not found.
 *  the original active workflow scheme does not have a draft.
