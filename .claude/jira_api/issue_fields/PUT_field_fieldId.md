# PUT /rest/api/3/field/{fieldId}
**operationId:** `updateCustomField`
**Summary:** Update custom field

Updates a custom field.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string
  - `searcherKey`: string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
