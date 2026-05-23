# GET /rest/api/3/screens/tabs
**operationId:** `getBulkScreenTabs`
**Summary:** Get bulk screen tabs

Returns the list of tabs for a bulk of screens.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [query] []integer(int64) — The list of screen IDs. To include multiple screen IDs, provide an ampersand-separated list. For example, `screenId=1000
- `tabId` [query] []integer(int64) — The list of tab IDs. To include multiple tab IDs, provide an ampersand-separated list. For example, `tabId=10000&tabId=1
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResult` [query] integer(int32) — The maximum number of items to return per page. The maximum number is 100,

## Responses
- 200: any
- 400: Returned if the screen ID or the tab ID is empty.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
