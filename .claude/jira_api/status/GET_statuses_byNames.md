# GET /rest/api/3/statuses/byNames
**operationId:** `getStatusesByName`
**Summary:** Bulk get statuses by name

Returns a list of the statuses specified by one or more status names.

**[Permissions](#permissions) required:**

 *  *Administer projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Administer Jira* [project permission.](https://confluence.atlassian.com/x/yodKLg)
 *  *Browse projects* [project permission.](https://confluence.atlassian.com/x/yodKLg)

## Parameters
- `name` [query] (required) []string — The list of status names. To include multiple names, provide an ampersand-separated list. For example, name=nameXX&name=
- `projectId` [query] string — The project the status is part of or null for global statuses.

## Responses
- 200: []object:
  - `description`: string
  - `id`: string
  - `name`: string
  - `scope`: StatusScope
  - `statusCategory`: string
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
