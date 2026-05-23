# DELETE /rest/api/3/role/{id}/actors
**operationId:** `deleteProjectRoleActorsFromRole`
**Summary:** Delete default actors from project role

Deletes the [default actors](#api-rest-api-3-resolution-get) from a project role. You may delete a group or user, but you cannot delete a group and a user in the same request.

Changing a project role's default actors does not affect project role members for projects already created.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
- `user` [query] string — The user account ID of the user to remove as a default actor.
- `groupId` [query] string — The group ID of the group to be removed as a default actor. This parameter cannot be used with the `group` parameter.
- `group` [query] string — The group name of the group to be removed as a default actor.This parameter cannot be used with the `groupId` parameter.

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
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have administrative permissions.
- 404: Returned if the project role is not found.
