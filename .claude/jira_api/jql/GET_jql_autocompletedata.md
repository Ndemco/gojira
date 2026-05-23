# GET /rest/api/3/jql/autocompletedata
**operationId:** `getAutoComplete`
**Summary:** Get field reference data (GET)

Returns reference data for JQL searches. This is a downloadable version of the documentation provided in [Advanced searching - fields reference](https://confluence.atlassian.com/x/gwORLQ) and [Advanced searching - functions reference](https://confluence.atlassian.com/x/hgORLQ), along with a list of JQL-reserved words. Use this information to assist with the programmatic creation of JQL queries or the validation of queries built in a custom query builder.

To filter visible field details by proje

## Responses
- 200: object:
  - `jqlReservedWords`: []string
  - `visibleFieldNames`: []FieldReferenceData
  - `visibleFunctionNames`: []FunctionReferenceData
- 401: Returned if the authentication credentials are incorrect.
