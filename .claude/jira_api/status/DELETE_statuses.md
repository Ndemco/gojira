# DELETE /rest/api/3/statuses
**operationId:** `deleteStatusesById`
**Summary:** Bulk delete Statuses

Deletes statuses by ID.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Administer Jira* [project permission.](https://confluence.atlassian.com/x/yodKLg)

## Parameters
- `id` [query] (required) []string — The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id=10000&id=10001.

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
