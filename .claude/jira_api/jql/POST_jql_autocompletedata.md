# POST /rest/api/3/jql/autocompletedata
**operationId:** `getAutoCompletePost`
**Summary:** Get field reference data (POST)

Returns reference data for JQL searches. This is a downloadable version of the documentation provided in [Advanced searching - fields reference](https://confluence.atlassian.com/x/gwORLQ) and [Advanced searching - functions reference](https://confluence.atlassian.com/x/hgORLQ), along with a list of JQL-reserved words. Use this information to assist with the programmatic creation of JQL queries or the validation of queries built in a custom query builder.

This operation can filter the custom fie

## Request Body
Content-Type: `application/json`
object:
  - `includeCollapsedFields`: boolean
  - `projectIds`: []integer(int64)

## Responses
- 200: object:
  - `jqlReservedWords`: []string
  - `visibleFieldNames`: []FieldReferenceData
  - `visibleFunctionNames`: []FunctionReferenceData
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect.
