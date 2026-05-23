# POST /rest/api/3/workflow/transitions/{transitionId}/properties
**operationId:** `createWorkflowTransitionProperty`
**Summary:** Create workflow transition property

This will be removed on [June 1, 2026](https://developer.atlassian.com/cloud/jira/platform/changelog/#CHANGE-2570); add transition properties using [Bulk update workflows](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflows-update-post) instead.

Adds a property to a workflow transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/z

## Parameters
- `transitionId` [path] (required) integer(int64) — The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next
- `key` [query] (required) string — The key of the property being added, also known as the name of the property. Set this to the same value as the `key` def
- `workflowName` [query] (required) string — The name of the workflow that the transition belongs to.
- `workflowMode` [query] string — The workflow status. Set to *live* for inactive workflows or *draft* for draft workflows. Active workflows cannot be edi

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
- 400: Returned if a workflow property with the same key is present on the transition.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow transition is not found.
