# DELETE /rest/api/3/projectCategory/{id}
**operationId:** `removeProjectCategory`
**Summary:** Delete project category

Deletes a project category.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — ID of the project category to delete.

## Responses
- 204: Returned if the request is successful.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the project category is not found.
