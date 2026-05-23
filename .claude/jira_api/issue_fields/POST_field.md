# POST /rest/api/3/field
**operationId:** `createCustomField`
**Summary:** Create custom field

Creates a custom field.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string
  - `searcherKey`: string
  - `type` (required): string

## Responses
- 201: object:
  - `clauseNames`: []string
  - `custom`: boolean
  - `id`: string
  - `key`: string
  - `name`: string
  - `navigable`: boolean
  - `orderable`: boolean
  - `schema`: allOf(JsonTypeBean)
  - `scope`: allOf(Scope)
  - `searchable`: boolean
- 400: Returned if:

 *  the user does not have permission to create custom fields.
 *  any of the request object properties have invalid or missing values.
