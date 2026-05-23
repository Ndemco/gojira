# GET /rest/api/3/issuetype/{issueTypeId}/properties
**operationId:** `getIssueTypePropertyKeys`
**Summary:** Get issue type property keys

Returns all the [issue type property](https://developer.atlassian.com/cloud/jira/platform/storing-data-without-a-database/#a-id-jira-entity-properties-a-jira-entity-properties) keys of the issue type.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) to get the property keys of any issue type.
 *  *Browse projects* [project permission](https://confluence.atlassian.com/x/yod

## Parameters
- `issueTypeId` [path] (required) string — The ID of the issue type.

## Responses
- 200: object:
  - `keys`: []PropertyKey
- 400: Returned if the issue type ID is invalid.
- 404: Returned if:

 *  the issue type is not found.
 *  the user does not have the required permissions.
