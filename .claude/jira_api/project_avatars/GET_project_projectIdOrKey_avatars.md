# GET /rest/api/3/project/{projectIdOrKey}/avatars
**operationId:** `getAllProjectAvatars`
**Summary:** Get all project avatars

Returns all project avatars, grouped by system and custom avatars.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or (case-sensitive) key of the project.

## Responses
- 200: object:
  - `custom`: []Avatar
  - `system`: []Avatar
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have permission to view the project.
