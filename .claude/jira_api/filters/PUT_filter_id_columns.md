# PUT /rest/api/3/filter/{id}/columns
**operationId:** `setColumns`
**Summary:** Set columns

Sets the columns for a filter. Only navigable fields can be set as columns. Use [Get fields](#api-rest-api-3-field-get) to get the list fields in Jira. A navigable field has `navigable` set to `true`.

The parameters for this resource are expressed as HTML form data. For example, in curl:

`curl -X PUT -d columns=summary -d columns=description https://your-domain.atlassian.net/rest/api/3/filter/10000/columns`

**[Permissions](#permissions) required:** Permission to access Jira, however, columns 

## Parameters
- `id` [path] (required) integer(int64) — The ID of the filter.

## Request Body
Content-Type: `*/*`
object:
  - `columns`: []string
Content-Type: `application/json`
object:
  - `columns`: []string
Content-Type: `multipart/form-data`
object:
  - `columns`: []string

## Responses
- 200: any
- 400: Returned if:

 *  a non-navigable field is set as a column.
 *  the user does not have permission to view the filter.
- 403: Returned if the requesting user is not an owner of the filter.
