# GET /rest/api/3/securitylevel/{id}
**operationId:** `getIssueSecurityLevel`
**Summary:** Get issue security level

Returns details of an issue security level.

Use [Get issue security scheme](#api-rest-api-3-issuesecurityschemes-id-get) to obtain the IDs of issue security levels associated with the issue security scheme.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** None.

## Parameters
- `id` [path] (required) string — The ID of the issue security level.

## Responses
- 200: object:
  - `description`: string
  - `id`: string
  - `isDefault`: boolean
  - `issueSecuritySchemeId`: string
  - `name`: string
  - `self`: string
- 401: Returned if the authentication credentials are incorrect.
- 404: Returned if the issue security level is not found.
