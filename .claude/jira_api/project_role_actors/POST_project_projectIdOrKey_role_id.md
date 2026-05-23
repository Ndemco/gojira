# POST /rest/api/3/project/{projectIdOrKey}/role/{id}
**operationId:** `addActorUsers`
**Summary:** Add actors to project role

Adds actors to a project role for the project.

To replace all actors for the project, use [Set actors for project role](#api-rest-api-3-project-projectIdOrKey-role-id-put).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project or *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `id` [path] (required) integer(int64) — The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.

## Request Body
Content-Type: `application/json`
object:
  - `group`: []string
  - `groupId`: []string
  - `user`: []string

## Responses
- 200: object:
  - `actors`: []RoleActor
  - `admin`: boolean
  - `currentUserRole`: boolean
  - `default`: boolean
  - `description`: string
  - `id`: integer(int64)
  - `name`: string
  - `roleConfigurable`: boolean
  - `scope`: allOf(Scope)
  - `self`: string(uri)
  - `translatedName`: string
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing or if the calling user lacks administrative permissions for the project.
- 404: Returned if:

 *  the project is not found.
 *  the user or group is not found.
 *  the group or user is not active.
