# DELETE /rest/api/3/version/{id}
**operationId:** `deleteVersion`
**Summary:** Delete version

Deletes a project version.

Deprecated, use [ Delete and replace version](#api-rest-api-3-version-id-removeAndSwap-post) that supports swapping version values in custom fields, in addition to the swapping for `fixVersion` and `affectedVersion` provided in this resource.

Alternative versions can be provided to update issues that use the deleted version in `fixVersion` or `affectedVersion`. If alternatives are not provided, occurrences of `fixVersion` and `affectedVersion` that contain the delete

## Parameters
- `id` [path] (required) string — The ID of the version.
- `moveFixIssuesTo` [query] string — The ID of the version to update `fixVersion` to when the field contains the deleted version. The replacement version mus
- `moveAffectedIssuesTo` [query] string — The ID of the version to update `affectedVersion` to when the field contains the deleted version. The replacement versio

## Responses
- 204: Returned if the version is deleted.
- 400: Returned if the request is invalid.
- 401: Returned if:

 *  the authentication credentials are incorrect.
 *  the user does not have the required permissions.
- 404: Returned if the version is not found.
