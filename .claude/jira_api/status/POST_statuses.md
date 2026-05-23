# POST /rest/api/3/statuses
**operationId:** `createStatuses`
**Summary:** Bulk create statuses

Creates statuses for a global or project scope.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Administer Jira* [project permission.](https://confluence.atlassian.com/x/yodKLg)

## Request Body
Content-Type: `application/json`
object:
  - `scope` (required): StatusScope
  - `statuses` (required): []StatusCreate

## Responses
- 200: []object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `scope`: StatusScope
  - `statusCategory`: string
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
- 409: Returned if another workflow configuration update task is ongoing.
