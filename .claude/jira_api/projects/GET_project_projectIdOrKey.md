# GET /rest/api/3/project/{projectIdOrKey}
**operationId:** `getProject`
**Summary:** Get project

Returns the [project details](https://confluence.atlassian.com/x/ahLpNw) for a project.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `properties` [query] []string — A list of project properties to return for the project. This parameter accepts a comma-separated list.

## Responses
- 200: object:
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
- 404: Returned if the project is not found or the user does not have permission to view it.
