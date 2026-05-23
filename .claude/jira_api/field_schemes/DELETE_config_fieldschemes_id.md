# DELETE /rest/api/3/config/fieldschemes/{id}
**operationId:** `deleteFieldAssociationScheme`
**Summary:** Delete a field scheme

Delete a specified field association scheme

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the field association scheme to delete.

## Responses
- 200: object:
  - `deleted`: boolean
  - `id`: string
- 400: any
- 401: any
- 403: any
- 404: any
- 409: any
