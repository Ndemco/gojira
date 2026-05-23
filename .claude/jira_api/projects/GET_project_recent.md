# GET /rest/api/3/project/recent
**operationId:** `getRecent`
**Summary:** Get recent projects

Returns a list of up to 20 projects recently viewed by the user that are still visible to the user.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** Projects are returned only where the user has one of:

 *  *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.
 *  *Administer Jira* [global permission](https

## Parameters
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `properties` [query] []StringList — EXPERIMENTAL. A list of project properties to return for the project. This parameter accepts a comma-separated list. Inv

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
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
