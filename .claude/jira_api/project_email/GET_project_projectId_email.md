# GET /rest/api/3/project/{projectId}/email
**operationId:** `getProjectEmail`
**Summary:** Get project's sender email

Returns the [project's sender email address](https://confluence.atlassian.com/x/dolKLg).

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectId` [path] (required) integer(int64) — The project ID.

## Responses
- 200: object:
  - `emailAddress`: string
  - `emailAddressStatus`: []string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to read project.
- 404: Returned if the project or project's sender email address is not found.
