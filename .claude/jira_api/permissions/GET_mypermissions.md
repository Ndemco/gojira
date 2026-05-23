# GET /rest/api/3/mypermissions
**operationId:** `getMyPermissions`
**Summary:** Get my permissions

Returns a list of permissions indicating which permissions the user has. Details of the user's permissions can be obtained in a global, project, issue or comment context.

The user is reported as having a project permission:

 *  in the global context, if the user has the project permission in any project.
 *  for a project, where the project permission is determined using issue data, if the user meets the permission's criteria for any issue in the project. Otherwise, if the user has the project

## Parameters
- `projectKey` [query] string — The key of project. Ignored if `projectId` is provided.
- `projectId` [query] string — The ID of project.
- `issueKey` [query] string — The key of the issue. Ignored if `issueId` is provided.
- `issueId` [query] string — The ID of the issue.
- `permissions` [query] string — A list of permission keys. (Required) This parameter accepts a comma-separated list. To get the list of available permis
- `projectUuid` [query] string — 
- `projectConfigurationUuid` [query] string — 
- `commentId` [query] string — The ID of the comment.

## Responses
- 200: object:
  - `permissions`: object
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
