# POST /rest/api/3/expression/analyse
**operationId:** `analyseExpression`
**Summary:** Analyse Jira expression

Analyses and validates Jira expressions.

As an experimental feature, this operation can also attempt to type-check the expressions.

Learn more about Jira expressions in the [documentation](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/).

**[Permissions](#permissions) required**: None.

## Parameters
- `check` [query] string — The check to perform:

## Request Body
Content-Type: `application/json`
object:
  - `contextVariables`: object
  - `expressions` (required): []string

## Responses
- 200: object:
  - `results` (required): []JiraExpressionAnalysis
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
