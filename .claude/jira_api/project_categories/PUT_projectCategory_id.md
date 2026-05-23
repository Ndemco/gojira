# PUT /rest/api/3/projectCategory/{id}
**operationId:** `updateProjectCategory`
**Summary:** Update project category

Updates a project category.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — 

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string(uri)

## Responses
- 200: object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `self`: string
- 400: Returned if:

 *  `name` has been modified and exceeds 255 characters.
 *  `description` has been modified and exceeds 1000 characters.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the project category is not found.
