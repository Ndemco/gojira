# POST /rest/api/3/config/fieldschemes
**operationId:** `createFieldAssociationScheme`
**Summary:** Create field scheme

Endpoint for creating a new field association scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string

## Responses
- 200: object:
  - `description`: string
  - `id`: integer(int64)
  - `links`: CreateFieldAssociationSchemeLinksBean
  - `name`: string
- 400: object
- 401: any
- 403: any
- 404: any
