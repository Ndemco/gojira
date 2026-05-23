# DELETE /rest/api/3/filter/{id}
**operationId:** `deleteFilter`
**Summary:** Delete filter

Delete a filter.

**[Permissions](#permissions) required:** Permission to access Jira, however filters can only be deleted by the creator of the filter or a user with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter to delete.

## Responses
- 204: Returned if the request is successful.
- 400: Returned if the filter is not found.
- 401: Returned if the user does not have permission to delete the filter.
