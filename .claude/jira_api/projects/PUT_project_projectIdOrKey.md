# PUT /rest/api/3/project/{projectIdOrKey}
**operationId:** `updateProject`
**Summary:** Update project

Updates the [project details](https://confluence.atlassian.com/x/ahLpNw) of a project.

All parameters are optional in the body of the request. Schemes will only be updated if they are included in the request, any omitted schemes will be left unchanged.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). is only needed when changing the schemes or project key. Otherwise you will only need *Administer Projects* [project perm

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Request Body
Content-Type: `application/json`
object:
  - `assigneeType`: string
  - `avatarId`: integer(int64)
  - `categoryId`: integer(int64)
  - `description`: string
  - `issueSecurityScheme`: integer(int64)
  - `key`: string
  - `lead`: string
  - `leadAccountId`: string
  - `name`: string
  - `notificationScheme`: integer(int64)
  - `permissionScheme`: integer(int64)
  - `releasedProjectKeys`: []string
  - `url`: string

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
- 403: Returned if:

 *  the user does not have the necessary permission to update project details.
 *  the permission scheme is being changed and the Jira instance is Jira Core Free or Jira Software Free. Permission schemes cannot be changed on free plans.
- 404: Returned if the project is not found.
