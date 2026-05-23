# GET /rest/api/3/workflow/search
**operationId:** `getWorkflowsPaginated`
**Summary:** Get workflows paginated

This will be removed on [June 1, 2026](https://developer.atlassian.com/cloud/jira/platform/changelog/#CHANGE-2569); use [Search workflows](#api-rest-api-3-workflows-search-get) instead.

Returns a [paginated](#pagination) list of published classic workflows. When workflow names are specified, details of those workflows are returned. Otherwise, all published classic workflows are returned.

This operation does not return next-gen workflows.

**[Permissions](#permissions) required:** *Administer J

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `workflowName` [query] []string — The name of a workflow to return. To include multiple workflows, provide an ampersand-separated list. For example, `work
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `queryString` [query] string — String used to perform a case-insensitive partial match with workflow name.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `isActive` [query] boolean — Filters active and inactive workflows.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Workflow
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
