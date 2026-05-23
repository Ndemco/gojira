# PUT /rest/api/3/config/fieldschemes/fields
**operationId:** `updateFieldsAssociatedWithSchemes`
**Summary:** Update fields associated with field schemes

Update fields associated with field association schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object

## Responses
- 200: object:
  - `results` (required): []FieldSchemeToFieldsPartialFailure
- 204: any
- 207: object:
  - `results` (required): []FieldSchemeToFieldsPartialFailure
- 400: object:
  - `results` (required): []FieldSchemeToFieldsPartialFailure
- 401: any
- 403: any
