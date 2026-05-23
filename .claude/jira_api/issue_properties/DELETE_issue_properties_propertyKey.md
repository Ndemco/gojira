# DELETE /rest/api/3/issue/properties/{propertyKey}
**operationId:** `bulkDeleteIssueProperty`
**Summary:** Bulk delete issue property

Deletes a property value from multiple issues. The issues to be updated can be specified by filter criteria.

The criteria the filter used to identify eligible issues are:

 *  `entityIds` Only issues from this list are eligible.
 *  `currentValue` Only issues with the property set to this value are eligible.

If both criteria is specified, they are joined with the logical *AND*: only issues that satisfy both criteria are considered eligible.

If no filter criteria are specified, all the issues 

## Parameters
- `propertyKey` [path] (required) string — The key of the property.

## Request Body
Content-Type: `application/json`
object:
  - `currentValue`: any
  - `entityIds`: []integer(int64)

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
