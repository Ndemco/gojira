# GET /rest/api/3/group/bulk
**operationId:** `bulkGetGroups`
**Summary:** Bulk get groups

Returns a [paginated](#pagination) list of groups.

**[Permissions](#permissions) required:** *Browse users and groups* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `groupId` [query] []string — The ID of a group. To specify multiple IDs, pass multiple `groupId` parameters. For example, `groupId=5b10a2844c20165700
- `groupName` [query] []string — The name of a group. To specify multiple names, pass multiple `groupName` parameters. For example, `groupName=administra
- `accessType` [query] string — The access level of a group. Valid values: 'site-admin', 'admin', 'user'.
- `applicationKey` [query] string — The application key of the product user groups to search for. Valid values: 'jira-servicedesk', 'jira-software', 'jira-p

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []GroupDetails
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 500: any
