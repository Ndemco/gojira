# POST /rest/api/3/config/fieldschemes/{id}/clone
**operationId:** `cloneFieldAssociationScheme`
**Summary:** Clone field scheme

Endpoint for cloning an existing field association scheme into a new one.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the source field association scheme to clone from

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
