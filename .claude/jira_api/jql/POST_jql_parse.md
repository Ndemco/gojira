# POST /rest/api/3/jql/parse
**operationId:** `parseJqlQueries`
**Summary:** Parse JQL query

Parses and validates JQL queries.

Validation is performed in context of the current user.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `validation` [query] (required) string — How to validate the JQL query and treat the validation results. Validation options include:

## Request Body
Content-Type: `application/json`
object:
  - `queries` (required): []string

## Responses
- 200: object:
  - `queries` (required): []ParsedJqlQuery
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect.
