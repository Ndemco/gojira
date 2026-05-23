# POST /rest/api/3/field/{fieldId}/context/{contextId}/project/remove
**operationId:** `removeCustomFieldContextFromProjects`
**Summary:** Remove custom field context from projects

Removes a custom field context from projects.

A custom field context without any projects applies to all projects. Removing all projects from a custom field context would result in it applying to all projects.

If any project in the request is not assigned to the context, or the operation would result in two global contexts for the field, the operation fails.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `fieldId` [path] (required) string — The ID of the custom field.
- `contextId` [path] (required) integer(int64) — The ID of the context.

## Request Body
Content-Type: `application/json`
object:
  - `projectIds` (required): []string

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any
