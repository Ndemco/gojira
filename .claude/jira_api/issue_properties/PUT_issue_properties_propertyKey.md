# PUT /rest/api/3/issue/properties/{propertyKey}
**operationId:** `bulkSetIssueProperty`
**Summary:** Bulk set issue property

Sets a property value on multiple issues.

The value set can be a constant or determined by a [Jira expression](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/). Expressions must be computable with constant complexity when applied to a set of issues. Expressions must also comply with the [restrictions](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#restrictions) that apply to all Jira expressions.

The issues to be updated can be specified by a filter

## Parameters
- `propertyKey` [path] (required) string — The key of the property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
object:
  - `expression`: string
  - `filter`: allOf(IssueFilterForBulkPropertySet)
  - `value`: any

## Responses
- 303: Returned if the request is successful.
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
