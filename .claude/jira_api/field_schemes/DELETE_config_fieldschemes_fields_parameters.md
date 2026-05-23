# DELETE /rest/api/3/config/fieldschemes/fields/parameters
**operationId:** `removeFieldAssociationSchemeItemParameters`
**Summary:** Remove field parameters

Remove field association parameters overrides for work types.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object

## Responses
- 200: Returned if the removal was successful.
- 204: any
- 207: object:
  - `results`: []SuccessOrErrorResults
- 400: object
- 401: any
- 403: any
