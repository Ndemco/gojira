# PUT /rest/api/3/issuetype/{issueTypeId}/properties/{propertyKey}
**operationId:** `setIssueTypeProperty`
**Summary:** Set issue type property

Creates or updates the value of the [issue type property](https://developer.atlassian.com/cloud/jira/platform/storing-data-without-a-database/#a-id-jira-entity-properties-a-jira-entity-properties). Use this resource to store and update data against an issue type.

The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://

## Parameters
- `issueTypeId` [path] (required) string — The ID of the issue type.
- `propertyKey` [path] (required) string — The key of the issue type property. The maximum length is 255 characters.

## Request Body
Content-Type: `application/json`
any

## Responses
- 200: any
- 201: any
- 400: Returned if:

 *  the issue type ID is invalid.
 *  a property value is not provided.
 *  the property value JSON content is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to modify the issue type.
- 404: Returned if:

 *  the issue type is not found.
 *  the user does not have the permission view the issue type.
