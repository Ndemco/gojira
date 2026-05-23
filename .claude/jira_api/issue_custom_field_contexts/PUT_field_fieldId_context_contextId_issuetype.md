# PUT /rest/api/3/field/{fieldId}/context/{contextId}/issuetype
**operationId:** `addIssueTypesToContext`
**Summary:** Add issue types to context

Adds issue types to a custom field context, appending the issue types to the issue types list.

A custom field context without any issue types applies to all issue types. Adding issue types to such a custom field context would result in it applying to only the listed issue types.

If any of the issue types exists in the custom field context, the operation fails and no issue types are added.

This API will not allow adding issue types to the global context from April 2026. Instead, an HTTP 400 re

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Request Body
Content-Type: `application/json`
object:
  - `issueTypeIds` (required): []string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
- 409: any
