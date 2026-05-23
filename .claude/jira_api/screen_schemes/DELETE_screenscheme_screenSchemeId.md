# DELETE /rest/api/3/screenscheme/{screenSchemeId}
**operationId:** `deleteScreenScheme`
**Summary:** Delete screen scheme

Deletes a screen scheme. A screen scheme cannot be deleted if it is used in an issue type screen scheme.

Only screens schemes used in classic projects can be deleted.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenSchemeId` [path] (required) string — The ID of the screen scheme.

## Responses
- 204: Returned if the screen scheme is deleted.
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
