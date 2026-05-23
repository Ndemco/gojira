# POST /rest/api/3/issue/properties
**operationId:** `bulkSetIssuesPropertiesList`
**Summary:** Bulk set issues properties by list

Sets or updates a list of entity property values on issues. A list of up to 10 entity properties can be specified along with up to 10,000 issues on which to set or update that list of entity properties.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON. The maximum length of single issue property value is 32768 characters. This operation can be accessed anonymously.

This operation is:

 *  transactional, either all properties are updated in all

## Request Body
Content-Type: `application/json`
object:
  - `entitiesIds`: []integer(int64)
  - `properties`: object

## Responses
- 303: Returned if the operation is successful.
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 409: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
