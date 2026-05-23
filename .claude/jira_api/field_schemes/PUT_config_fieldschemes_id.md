# PUT /rest/api/3/config/fieldschemes/{id}
**operationId:** `updateFieldAssociationScheme`
**Summary:** Update field scheme

Endpoint for updating an existing field association scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — 

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string

## Responses
- 200: object:
  - `description`: string
  - `id`: integer(int64)
  - `links`: UpdateFieldAssociationSchemeLinksBean
  - `name`: string
- 400: object
- 401: any
- 403: any
- 404: any
