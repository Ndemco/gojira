# DELETE /rest/api/3/issue/{issueIdOrKey}/properties/{propertyKey}
**operationId:** `deleteIssueProperty`
**Summary:** Delete issue property

Deletes an issue's property.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Edit issues* [project permissions](https://confluence.atlassian.com/x/yodKLg) for the project containing the issue.
 *  If [issue-level security](https://confluence.atlassian.com/x/J4lKLg) is configured, issue-level security permission to view the issue.

## Parameters
- `issueIdOrKey` [path] (required) string — The key or ID of the issue.
- `propertyKey` [path] (required) string — The key of the property.

## Responses
- 204: Returned if the request is successful.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue or property is not found, or the user does not have permission to edit the issue.
