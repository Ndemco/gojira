# POST /rest/api/3/project/{projectIdOrKey}/restore
**operationId:** `restore`
**Summary:** Restore deleted or archived project

Restores a project that has been archived or placed in the Jira recycle bin.

**[Permissions](#permissions) required:**

 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg)for Company managed projects.
 *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project for Team managed projects.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).

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
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have the necessary permission.
