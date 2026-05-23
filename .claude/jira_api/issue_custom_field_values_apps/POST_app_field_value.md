# POST /rest/api/3/app/field/value
**operationId:** `updateMultipleCustomFieldValues`
**Summary:** Update custom fields

Updates the value of one or more custom fields on one or more issues. Combinations of custom field and issue should be unique within the request.

Apps can only perform this operation on [custom fields](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/) and [custom field types](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field-type/) declared in their own manifests.

**[Permissions](#permissions) required:** Onl

## Parameters
- `generateChangelog` [query] boolean — Whether to generate a changelog for this update.
- `generateAppEvents` [query] boolean — Whether to generate app events for this update. Suppresses Forge, Connect, OAuth 2.0, and admin-configured webhooks (reg

## Request Body
Content-Type: `application/json`
object:
  - `updates`: []MultipleCustomFieldValuesUpdate

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 403: Returned if the request is not authenticated as the app that provided all the fields.
- 404: Returned if any field is not found.
