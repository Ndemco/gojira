# PUT /rest/atlassian-connect/1/migration/field
**operationId:** `AppIssueFieldValueUpdateResource.updateIssueFields_put`
**Summary:** Bulk update custom field value

Updates the value of a custom field added by Connect apps on one or more issues.
The values of up to 200 custom fields can be updated.

**[Permissions](#permissions) required:** Only Connect apps can make this request

## Parameters
- `Atlassian-Transfer-Id` [header] (required) string(uuid) — The ID of the transfer.

## Request Body
Content-Type: `application/json`
object:
  - `updateValueList`: []ConnectCustomFieldValue

## Responses
- 200: any
- 400: Returned if the request is invalid.
- 403: Returned if:
* the transfer ID is not found.
* the authorisation credentials are incorrect or missing.
