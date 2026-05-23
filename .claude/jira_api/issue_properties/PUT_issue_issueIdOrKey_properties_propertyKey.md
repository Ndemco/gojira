# PUT /rest/api/3/issue/{issueIdOrKey}/properties/{propertyKey}
**operationId:** `setIssueProperty`
**Summary:** Set issue property

Sets the value of an issue's property. Use this resource to store custom data against an issue.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Browse projects* and *Edit issues* [project permissions](https://confluence.atlassian.com/x/yodKLg) for the project containing the issue.
 *  If [issue-level secur

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `propertyKey` [path] (required) string — The key of the issue property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to edit the issue.
- 404: Returned if the issue is not found or the user does not have permission to view the issue.
