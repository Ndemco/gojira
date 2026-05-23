# GET /rest/api/3/worklog/updated
**operationId:** `getIdsOfWorklogsModifiedSince`
**Summary:** Get IDs of updated worklogs

Returns a list of IDs and update timestamps for worklogs updated after a date and time.

This resource is paginated, with a limit of 1000 worklogs per page. Each page lists worklogs from oldest to youngest. If the number of items in the date range exceeds 1000, `until` indicates the timestamp of the youngest item on the page. Also, `nextPage` provides the URL for the next page of worklogs. The `lastPage` parameter is set to true on the last page of worklogs.

This resource does not return worklo

## Parameters
- `since` [query] integer(int64) — The date and time, as a UNIX timestamp in milliseconds, after which updated worklogs are returned.
- `expand` [query] string — Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `prope

## Responses
- 200: object:
  - `lastPage`: boolean
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `since`: integer(int64)
  - `until`: integer(int64)
  - `values`: []ChangedWorklog
- 401: Returned if the authentication credentials are incorrect or missing.
