# POST /rest/api/3/expression/eval
**operationId:** `evaluateJiraExpression`
**Summary:** Currently being removed. Evaluate Jira expression

Endpoint is currently being removed. [More details](https://developer.atlassian.com/changelog/#CHANGE-2046)

Evaluates a Jira expression and returns its value.

This resource can be used to test Jira expressions that you plan to use elsewhere, or to fetch data in a flexible way. Consult the [Jira expressions documentation](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/) for more details.

#### Context variables ####

The following context variables are available to Jira ex

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts `meta.complexity` tha

## Request Body
Content-Type: `application/json`
object:
  - `context`: allOf(JiraExpressionEvalContextBean)
  - `expression` (required): string

## Responses
- 200: object:
  - `meta`: allOf(JiraExpressionEvaluationMetaDataBean)
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
