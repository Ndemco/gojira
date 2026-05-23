# GET /rest/api/3/workflows/search
**operationId:** `searchWorkflows`
**Summary:** Search workflows

Returns a [paginated](#pagination) list of global and project workflows. If workflow names are specified in the query string, details of those workflows are returned. Otherwise, all workflows are returned.

**[Permissions](#permissions) required:**

 *  *Administer Jira* global permission to access all, including project-scoped, workflows
 *  At least one of the *Administer projects* and *View (read-only) workflow* project permissions to access project-scoped workflows

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `queryString` [query] string — String used to perform a case-insensitive partial match with workflow name.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `scope` [query] string — The scope of the workflow. Global for company-managed projects and Project for team-managed projects.
- `isActive` [query] boolean — Filters active and inactive workflows.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string
  - `self`: string
  - `startAt`: integer(int64)
  - `statuses`: []JiraWorkflowStatus
  - `total`: integer(int64)
  - `values`: []JiraWorkflow
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing, or the caller doesn't have permissions to perform the operation.
