# PUT /rest/api/3/app/field/{fieldIdOrKey}/context/configuration
**operationId:** `updateCustomFieldConfiguration`
**Summary:** Update custom field configurations

Update the configuration for contexts of a custom field of a [type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field-type/) created by a [Forge app](https://developer.atlassian.com/platform/forge/).

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). Jira permissions are not required for the Forge app that created the custom field type.

## Parameters
- `fieldIdOrKey` [path] (required) string — The ID or key of the custom field, for example `customfield_10000`.

## Request Body
Content-Type: `application/json`
object:
  - `configurations` (required): []ContextualConfiguration

## Responses
- 200: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user is not a Jira admin or the request is not authenticated as from the app that provided the field.
- 404: Returned if the custom field is not found.
