# GET /rest/api/3/project
**operationId:** `getAllProjects`
**Summary:** Get all projects

Returns all projects visible to the user. Deprecated, use [ Get projects paginated](#api-rest-api-3-project-search-get) that supports search and pagination.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Projects are returned only where the user has *Browse Projects* or *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `recent` [query] integer(int32) — Returns the user's most recently accessed projects. You may specify the number of results to return up to a maximum of 2
- `properties` [query] []string — A list of project properties to return for the project. This parameter accepts a comma-separated list.

## Responses
- 200: []object:
  - `archived`: boolean
  - `archivedBy`: allOf(User)
  - `archivedDate`: string(date-time)
  - `assigneeType`: string
  - `avatarUrls`: allOf(AvatarUrlsBean)
  - `components`: []ProjectComponent
  - `deleted`: boolean
  - `deletedBy`: allOf(User)
  - `deletedDate`: string(date-time)
  - `description`: string
  - `email`: string
  - `expand`: string
  - `favourite`: boolean
  - `id`: string
  - `insight`: allOf(ProjectInsight)
  - `isPrivate`: boolean
  - `issueTypeHierarchy`: allOf(Hierarchy)
  - `issueTypes`: []IssueTypeDetails
  - `key`: string
  - `landingPageInfo`: allOf(ProjectLandingPageInfo)
  - ... (14 more fields)
- 401: Returned if the authentication credentials are incorrect or missing.
