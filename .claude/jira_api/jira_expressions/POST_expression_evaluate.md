# POST /rest/api/3/expression/evaluate
**operationId:** `evaluateJSISJiraExpression`
**Summary:** Evaluate Jira expression using enhanced search API

Evaluates a Jira expression and returns its value. The difference between this and `eval` is that this endpoint uses the enhanced search API when evaluating JQL queries. This API is eventually consistent, unlike the strongly consistent `eval` API. This allows for better performance and scalability. In addition, this API's response for JQL evaluation is based on a scrolling view (backed by a `nextPageToken`) instead of a paginated view (backed by `startAt` and `totalCount`).

This resource can be

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts `meta.complexity` tha

## Request Body
Content-Type: `application/json`
object:
  - `context`: allOf(JiraExpressionEvaluateContextBean)
  - `expression` (required): string

## Responses
- 200: object:
  - `meta`: allOf(JExpEvaluateMetaDataBean)
  - `value` (required): any
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
