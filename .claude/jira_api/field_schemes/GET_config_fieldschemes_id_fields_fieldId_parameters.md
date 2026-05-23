# GET /rest/api/3/config/fieldschemes/{id}/fields/{fieldId}/parameters
**operationId:** `getFieldAssociationSchemeItemParameters`
**Summary:** Get field parameters

Retrieve field association parameters on a field association scheme

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — the ID of the field association scheme to retrieve parameters for
- `fieldId` [path] (required) string — the ID of the field

## Responses
- 200: object:
  - `fieldId` (required): string
  - `parameters`: FieldAssociationParameters
  - `workTypeParameters`: []WorkTypeParameters
- 400: object
- 401: any
- 403: any
