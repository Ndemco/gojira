# POST /rest/api/3/projectCategory
**operationId:** `createProjectCategory`
**Summary:** Create project category

Creates a project category.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)

## Responses
- 201: object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)
- 400: Returned if:

 *  `name` is not provided or exceeds 255 characters.
 *  `description` exceeds 1000 characters.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 409: Returned if the project category name is in use.
