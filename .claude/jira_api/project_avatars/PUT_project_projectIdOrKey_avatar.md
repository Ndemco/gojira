# PUT /rest/api/3/project/{projectIdOrKey}/avatar
**operationId:** `updateProjectAvatar`
**Summary:** Set project avatar

Sets the avatar displayed for a project.

Use [Load project avatar](#api-rest-api-3-project-projectIdOrKey-avatar2-post) to store avatars against the project, before using this operation to set the displayed avatar.

**[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

## Parameters
- `projectIdOrKey` [path] (required) string — The ID or (case-sensitive) key of the project.

## Request Body
Content-Type: `application/json`
object:
  - `fileName`: string
  - `id` (required): string
  - `isDeletable`: boolean
  - `isSelected`: boolean
  - `isSystemAvatar`: boolean
  - `owner`: string
  - `urls`: object

## Responses
- 204: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to administer the project.
- 404: Returned if the project or avatar is not found or the user does not have permission to view the project.
