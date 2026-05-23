# POST /rest/api/3/changelog/bulkfetch
**operationId:** `getBulkChangelogs`
**Summary:** Bulk fetch changelogs

Bulk fetch changelogs for multiple issues and filter by fields

Returns a paginated list of all changelogs for given issues sorted by changelog date and issue IDs, starting from the oldest changelog and smallest issue ID.

Issues are identified by their ID or key, and optionally changelogs can be filtered by their field IDs. You can request the changelogs of up to 1000 issues and can filter them by up to 10 field IDs.

**[Permissions](#permissions) required:**

 *  *Browse projects* [project per

## Request Body
Content-Type: `application/json`
object:
  - `fieldIds`: []string
  - `issueIdsOrKeys` (required): []string
  - `maxResults`: integer(int32)
  - `nextPageToken`: string

## Responses
- 200: object:
  - `issueChangeLogs`: []IssueChangeLog
  - `nextPageToken`: string
- 400: Returned if there are input validation problems such as no issue IDs/keys were present, or more than 1000 issue IDs/keys were requested.
