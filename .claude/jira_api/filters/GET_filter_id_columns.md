# GET /rest/api/3/filter/{id}/columns
**operationId:** `getColumns`
**Summary:** Get columns

Returns the columns configured for a filter. The column configuration is used when the filter's results are viewed in *List View* with the *Columns* set to *Filter*.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None, however, column details are only returned for:

 *  filters owned by the user.
 *  filters shared with a group that the user is a member of.
 *  filters shared with a private project that the user has *Browse projects* [project permission](

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.

## Responses
- 200: []object:
  - `label`: string
  - `value`: string
- 400: Returned if the user does not have permission to view the filter.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if a column configuration is not set for the filter.
