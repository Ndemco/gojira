# PUT /rest/api/3/workflow/transitions/{transitionId}/properties
**operationId:** `updateWorkflowTransitionProperty`
**Summary:** Update workflow transition property

This will be removed on [June 1, 2026](https://developer.atlassian.com/cloud/jira/platform/changelog/#CHANGE-2570); update transition properties using [Bulk update workflows](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflows-update-post) instead.

Updates a workflow transition by changing the property value. Trying to update a property that does not exist results in a new property being added to the transition. Transition properties are use

## Parameters
- `transitionId` [path] (required) integer(int64) — The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next
- `key` [query] (required) string — The key of the property being updated, also known as the name of the property. Set this to the same value as the `key` d
- `workflowName` [query] (required) string — The name of the workflow that the transition belongs to.
- `workflowMode` [query] string — The workflow status. Set to `live` for inactive workflows or `draft` for draft workflows. Active workflows cannot be edi

## Request Body
Content-Type: `application/json`
object:
  - `id`: string
  - `key`: string
  - `value` (required): string

## Responses
- 200: object:
  - `id`: string
  - `key`: string
  - `value` (required): string
- 304: Returned if no changes were made by the request. For example, attempting to update a property with its current value.
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow transition is not found.
