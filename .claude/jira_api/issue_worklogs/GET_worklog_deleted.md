# GET /rest/api/3/worklog/deleted
**operationId:** `getIdsOfWorklogsDeletedSince`
**Summary:** Get IDs of deleted worklogs

Returns a list of IDs and delete timestamps for worklogs deleted after a date and time.

This resource is paginated, with a limit of 1000 worklogs per page. Each page lists worklogs from oldest to youngest. If the number of items in the date range exceeds 1000, `until` indicates the timestamp of the youngest item on the page. Also, `nextPage` provides the URL for the next page of worklogs. The `lastPage` parameter is set to true on the last page of worklogs.

This resource does not return worklo

## Parameters
- `since` [query] integer(int64) — The date and time, as a UNIX timestamp in milliseconds, after which deleted worklogs are returned.

## Responses
- 200: object:
  - `lastPage`: boolean
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `since`: integer(int64)
  - `until`: integer(int64)
  - `values`: []ChangedWorklog
- 401: Returned if the authentication credentials are incorrect or missing.
