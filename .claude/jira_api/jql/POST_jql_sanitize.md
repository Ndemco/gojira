# POST /rest/api/3/jql/sanitize
**operationId:** `sanitiseJqlQueries`
**Summary:** Sanitize JQL queries

Sanitizes one or more JQL queries by converting readable details into IDs where a user doesn't have permission to view the entity.

For example, if the query contains the clause *project = 'Secret project'*, and a user does not have browse permission for the project "Secret project", the sanitized query replaces the clause with *project = 12345"* (where 12345 is the ID of the project). If a user has the required permission, the clause is not sanitized. If the account ID is null, sanitizing is pe

## Request Body
Content-Type: `application/json`
object:
  - `queries` (required): []JqlQueryToSanitize

## Responses
- 200: object:
  - `queries`: []SanitizedJqlQuery
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
