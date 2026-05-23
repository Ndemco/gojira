# GET /rest/api/3/issuetype/{issueTypeId}/properties/{propertyKey}
**operationId:** `getIssueTypeProperty`
**Summary:** Get issue type property

Returns the key and value of the [issue type property](https://developer.atlassian.com/cloud/jira/platform/storing-data-without-a-database/#a-id-jira-entity-properties-a-jira-entity-properties).

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) to get the details of any issue type.
 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) to get 

## Parameters
- `issueTypeId` [path] (required) string — The ID of the issue type.
- `propertyKey` [path] (required) string — The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get

## Responses
- 200: object:
  - `key`: string
  - `value`: any
- 400: Returned if the issue type ID is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue type or property is not found or the user does not have the required permissions.
