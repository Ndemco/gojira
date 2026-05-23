# PUT /rest/api/3/app/field/{fieldIdOrKey}/value
**operationId:** `updateCustomFieldValue`
**Summary:** Update custom field value

Updates the value of a custom field on one or more issues.

Apps can only perform this operation on [custom fields](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/) and [custom field types](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field-type/) declared in their own manifests.

**[Permissions](#permissions) required:** Only the app that owns the custom field or custom field type can update its values with th

## Parameters
- `fieldIdOrKey` [path] (required) string — The ID or key of the custom field. For example, `customfield_10010`.
- `generateChangelog` [query] boolean — Whether to generate a changelog for this update.
- `generateAppEvents` [query] boolean — Whether to generate app events for this update. Suppresses Forge, Connect, OAuth 2.0, and admin-configured webhooks (reg

## Request Body
Content-Type: `application/json`
object:
  - `updates`: []CustomFieldValueUpdate

## Responses
- 204: any
- 400: Returned if the request is invalid.
- 403: Returned if the request is not authenticated as the app that provided the field.
- 404: Returned if the field is not found.
