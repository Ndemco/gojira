# GET /rest/api/3/config/fieldschemes/{id}
**operationId:** `getFieldAssociationSchemeById`
**Summary:** Get field scheme

Endpoint for fetching a field association scheme by its ID

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The scheme id to fetch

## Responses
- 200: object:
  - `description`: string
  - `id`: string
  - `isDefault`: boolean
  - `links`: FieldAssociationSchemeLinks
  - `name`: string
- 403: any
- 404: any
