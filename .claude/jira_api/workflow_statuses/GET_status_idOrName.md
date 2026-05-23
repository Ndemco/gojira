# GET /rest/api/3/status/{idOrName}
**operationId:** `getStatus`
**Summary:** Get status

Returns a status. The status must be associated with an active workflow to be returned.

If a name is used on more than one status, only the status found first is returned. Therefore, identifying the status by its ID may be preferable.

This operation can be accessed anonymously.

[Permissions](#permissions) required: *Browse projects* [project permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-project-permissions/) for the project.

## Parameters
- `idOrName` [path] (required) string — The ID or name of the status.

## Responses
- 200: object:
  - `description`: string
  - `iconUrl`: string
  - `id`: string
  - `name`: string
  - `scope`: allOf(Scope)
  - `self`: string
  - `statusCategory`: allOf(StatusCategory)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the status is not found.
 *  the status is not associated with a workflow.
 *  the user does not have the required permissions.
