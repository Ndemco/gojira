# POST /rest/api/3/role
**operationId:** `createProjectRole`
**Summary:** Create project role

Creates a new project role with no [default actors](#api-rest-api-3-resolution-get). You can use the [Add default actors to project role](#api-rest-api-3-role-id-actors-post) operation to add default actors to the project role after creating it.

*Note that although a new project role is available to all projects upon creation, any default actors that are associated with the project role are not added to projects that existed prior to the role being created.*<

**[Permissions](#permissions) requ

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string

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
- 400: Returned if the request is not valid. The `name` cannot be empty or start or end with whitespace.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have administrative permissions.
- 409: Returned if a project role with the provided name already exists.
