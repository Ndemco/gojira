# GET /rest/api/3/status
**operationId:** `getStatuses`
**Summary:** Get all statuses

Returns a list of all statuses associated with active workflows.

This operation can be accessed anonymously.

[Permissions](#permissions) required: *Browse projects* [project permission](https://support.atlassian.com/jira-cloud-administration/docs/manage-project-permissions/) for the project.

## Responses
- 200: []object:
  - `description`: string
  - `iconUrl`: string
  - `id`: string
  - `name`: string
  - `scope`: allOf(Scope)
  - `self`: string
  - `statusCategory`: allOf(StatusCategory)
- 401: Returned if the authentication credentials are incorrect or missing.
