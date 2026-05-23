# PUT /rest/api/3/config/fieldschemes/fields/parameters
**operationId:** `updateFieldAssociationSchemeItemParameters`
**Summary:** Update field parameters

Update field association item parameters in field association schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object

## Responses
- 200: object:
  - `results` (required): []UpdateFieldSchemeParametersPartialFailure
- 204: any
- 207: object:
  - `results` (required): []UpdateFieldSchemeParametersPartialFailure
- 400: object:
  - `results` (required): []UpdateFieldSchemeParametersPartialFailure
- 401: any
- 403: any
