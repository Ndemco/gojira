# POST /rest/api/3/project
**operationId:** `createProject`
**Summary:** Create project

Creates a project based on a project type template, as shown in the following table:

| Project Type Key | Project Template Key |  
|--|--|  
| `business` | `com.atlassian.jira-core-project-templates:jira-core-simplified-content-management`, `com.atlassian.jira-core-project-templates:jira-core-simplified-document-approval`, `com.atlassian.jira-core-project-templates:jira-core-simplified-lead-tracking`, `com.atlassian.jira-core-project-templates:jira-core-simplified-process-control`, `com.atlassi

## Request Body
Content-Type: `application/json`
object:
  - `assigneeType`: string
  - `avatarId`: integer(int64)
  - `categoryId`: integer(int64)
  - `description`: string
  - `fieldConfigurationScheme`: integer(int64)
  - `fieldScheme`: integer(int64)
  - `issueSecurityScheme`: integer(int64)
  - `issueTypeScheme`: integer(int64)
  - `issueTypeScreenScheme`: integer(int64)
  - `key` (required): string
  - `lead`: string
  - `leadAccountId`: string
  - `name` (required): string
  - `notificationScheme`: integer(int64)
  - `permissionScheme`: integer(int64)
  - `projectTemplateKey`: string
  - `projectTypeKey`: string
  - `url`: string
  - `workflowScheme`: integer(int64)

## Responses
- 201: object:
  - `id` (required): integer(int64)
  - `key` (required): string
  - `self` (required): string(uri)
- 400: Returned if the request is not valid and the project could not be created.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have permission to create projects.
