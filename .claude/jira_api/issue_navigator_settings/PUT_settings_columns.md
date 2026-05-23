# PUT /rest/api/3/settings/columns
**operationId:** `setIssueNavigatorDefaultColumns`
**Summary:** Set issue navigator default columns

Sets the default issue navigator columns.

The `columns` parameter accepts a navigable field value and is expressed as HTML form data. To specify multiple columns, pass multiple `columns` parameters. For example, in curl:

`curl -X PUT -d columns=summary -d columns=description https://your-domain.atlassian.net/rest/api/3/settings/columns`

If no column details are sent, then all default columns are removed.

A navigable field is one that can be used as a column on the issue navigator. Find detai

## Request Body
Content-Type: `*/*`
object:
  - `columns`: []string
Content-Type: `multipart/form-data`
object:
  - `columns`: []string

## Responses
- 200: Returned if the request is successful.
- 400: Returned if invalid parameters are passed.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if a navigable field value is not found.
