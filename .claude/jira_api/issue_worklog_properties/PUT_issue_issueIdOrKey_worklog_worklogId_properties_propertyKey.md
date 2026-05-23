# PUT /rest/api/3/issue/{issueIdOrKey}/worklog/{worklogId}/properties/{propertyKey}
**operationId:** `setWorklogProperty`
**Summary:** Set worklog property

Sets the value of a worklog property. Use this operation to store custom data against the worklog.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project that the issue is in.
 *  If [issue-level security](https://con

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `worklogId` [path] (required) string — The ID of the worklog.
- `propertyKey` [path] (required) string — The key of the issue property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
- 400: Returned if the worklog ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to edit the worklog.
- 404: Returned if:

 *  the issue or worklog is not found.
 *  the user does not have permission to view the issue or worklog.
