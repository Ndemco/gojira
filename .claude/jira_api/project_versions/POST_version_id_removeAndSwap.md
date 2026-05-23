# POST /rest/api/3/version/{id}/removeAndSwap
**operationId:** `deleteAndReplaceVersion`
**Summary:** Delete and replace version

Deletes a project version.

Alternative versions can be provided to update issues that use the deleted version in `fixVersion`, `affectedVersion`, or any version picker custom fields. If alternatives are not provided, occurrences of `fixVersion`, `affectedVersion`, and any version picker custom field, that contain the deleted version, are cleared. Any replacement version must be in the same project as the version being deleted and cannot be the version being deleted.

This operation can be acces

## Parameters
- `id` [path] (required) string — The ID of the version.

## Request Body
Content-Type: `application/json`
object:
  - `customFieldReplacementList`: []CustomFieldReplacement
  - `moveAffectedIssuesTo`: integer(int64)
  - `moveFixIssuesTo`: integer(int64)

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the version to delete is not found.
 *  the user does not have the required permissions.
