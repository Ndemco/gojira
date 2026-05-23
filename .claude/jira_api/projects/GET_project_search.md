# GET /rest/api/3/project/search
**operationId:** `searchProjects`
**Summary:** Get projects paginated

Returns a [paginated](#pagination) list of projects visible to the user.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Projects are returned only where the user has one of:

 *  *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer Jira* [global permission](https://confluence.atlassian.com

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page. Must be less than or equal to 100. If a value greater than 100 is provid
- `orderBy` [query] string — [Order](#ordering) the results by a field.
- `id` [query] []integer(int64) — The project IDs to filter the results by. To include multiple IDs, provide an ampersand-separated list. For example, `id
- `keys` [query] []string — The project keys to filter the results by. To include multiple keys, provide an ampersand-separated list. For example, `
- `query` [query] string — Filter the results using a literal string. Projects with a matching `key` or `name` are returned (case insensitive).
- `typeKey` [query] string — Orders results by the [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesa
- `categoryId` [query] integer(int64) — The ID of the project's category. A complete list of category IDs is found using the [Get all project categories](#api-r
- `action` [query] string — Filter results by projects for which the user can:
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `status` [query] []string — EXPERIMENTAL. Filter results by project status:
- `properties` [query] []StringList — EXPERIMENTAL. A list of project properties to return for the project. This parameter accepts a comma-separated list.
- `propertyQuery` [query] string — EXPERIMENTAL. A query string used to search properties. The query string cannot be specified using a JSON object. For ex

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Project
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if no projects matching the search criteria are found.
