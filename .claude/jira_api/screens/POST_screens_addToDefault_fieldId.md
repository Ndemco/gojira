# POST /rest/api/3/screens/addToDefault/{fieldId}
**operationId:** `addFieldToDefaultScreen`
**Summary:** Add field to default screen

Adds a field to the default tab of the default screen.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the field.

## Responses
- 200: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the field it not found or the field is already present.
