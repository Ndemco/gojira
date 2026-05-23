# GET /rest/api/3/project/{projectId}/hierarchy
**operationId:** `getHierarchy`
**Summary:** Get project issue type hierarchy

Get the issue type hierarchy for a next-gen project.

The issue type hierarchy for a project consists of:

 *  *Epic* at level 1 (optional).
 *  One or more issue types at level 0 such as *Story*, *Task*, or *Bug*. Where the issue type *Epic* is defined, these issue types are used to break down the content of an epic.
 *  *Subtask* at level -1 (optional). This issue type enables level 0 issue types to be broken down into components. Issues based on a level -1 issue type must have a parent issue.

## Parameters
- `projectId` [path] (required) integer(int64) — The ID of the project.

## Responses
- 200: object:
  - `hierarchy`: []ProjectIssueTypesHierarchyLevel
  - `projectId`: integer(int64)
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have the necessary permission.
