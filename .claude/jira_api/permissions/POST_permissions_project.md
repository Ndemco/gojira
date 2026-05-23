# POST /rest/api/3/permissions/project
**operationId:** `getPermittedProjects`
**Summary:** Get permitted projects

Returns all the projects where the user is granted a list of project permissions.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Request Body
Content-Type: `application/json`
object:
  - `permissions` (required): []string

## Responses
- 200: object:
  - `projects`: []ProjectIdentifierBean
- 400: Returned if a project permission is not found.
- 401: Returned if the authentication credentials are incorrect or missing.
