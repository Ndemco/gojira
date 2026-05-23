# GET /rest/api/3/statuses
**operationId:** `getStatusesById`
**Summary:** Bulk get statuses

Returns a list of the statuses specified by one or more status IDs.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Administer Jira* [project permission.](https://confluence.atlassian.com/x/yodKLg)

## Parameters
- `id` [query] (required) []string — The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id=10000&id=10001.

## Responses
- 200: []object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `scope`: StatusScope
  - `statusCategory`: string
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
