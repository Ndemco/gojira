# PUT /rest/api/3/field/{fieldId}/context/defaultValue
**operationId:** `setDefaultValues`
**Summary:** Set custom field contexts default values

Sets default for contexts of a custom field. Default are defined using these objects:

 *  `CustomFieldContextDefaultValueDate` (type `datepicker`) for date fields.
 *  `CustomFieldContextDefaultValueDateTime` (type `datetimepicker`) for date-time fields.
 *  `CustomFieldContextDefaultValueSingleOption` (type `option.single`) for single choice select lists and radio buttons.
 *  `CustomFieldContextDefaultValueMultipleOption` (type `option.multiple`) for multiple choice select lists and checkboxe

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.

## Request Body
Content-Type: `application/json`
object:
  - `defaultValues`: []CustomFieldContextDefaultValue

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
