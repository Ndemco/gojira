# PUT /rest/api/3/statuses
**operationId:** `updateStatuses`
**Summary:** Bulk update statuses

Updates statuses by ID.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Administer Jira* [project permission.](https://confluence.atlassian.com/x/yodKLg)

## Request Body
Content-Type: `application/json`
object:
  - `statuses` (required): []StatusUpdate

## Responses
- 204: any
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 409: Returned if another workflow configuration update task is ongoing.
