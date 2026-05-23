# PUT /rest/api/3/screenscheme/{screenSchemeId}
**operationId:** `updateScreenScheme`
**Summary:** Update screen scheme

Updates a screen scheme. Only screen schemes used in classic projects can be updated.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenSchemeId` [path] (required) string — The ID of the screen scheme.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string
  - `screens`: allOf(UpdateScreenTypes)

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
