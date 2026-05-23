# POST /rest/api/3/field/{fieldId}/context
**operationId:** `createCustomFieldContext`
**Summary:** Create custom field context

Creates a custom field context.

If `projectIds` is empty, a global context is created. A global context is one that applies to all project. If `issueTypeIds` is empty, the context applies to all issue types.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `id`: string
  - `issueTypeIds`: []string
  - `name` (required): string
  - `projectIds`: []string

## Responses
- 201: object:
  - `description`: string
  - `id`: string
  - `issueTypeIds`: []string
  - `name` (required): string
  - `projectIds`: []string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the field, project, or issue type is not found.
- 409: any
