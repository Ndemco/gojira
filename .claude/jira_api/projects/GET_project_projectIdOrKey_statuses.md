# GET /rest/api/3/project/{projectIdOrKey}/statuses
**operationId:** `getAllStatuses`
**Summary:** Get all statuses for project

Returns the valid statuses for a project. The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).

## Responses
- 200: []object:
  - `id` (required): string
  - `name` (required): string
  - `self` (required): string
  - `statuses` (required): []StatusDetails
  - `subtask` (required): boolean
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have permission to view it.
