# DELETE /rest/api/3/config/fieldschemes/fields
**operationId:** `removeFieldsAssociatedWithSchemes`
**Summary:** Remove fields associated with field schemes

Remove fields associated with field association schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object

## Responses
- 200: object:
  - `results` (required): []MinimalFieldSchemeToFieldsPartialFailure
- 204: any
- 207: object:
  - `results` (required): []MinimalFieldSchemeToFieldsPartialFailure
- 400: object:
  - `results` (required): []MinimalFieldSchemeToFieldsPartialFailure
- 401: any
- 403: any
