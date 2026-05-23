# GET /rest/api/3/workflow/transitions/{transitionId}/properties
**operationId:** `getWorkflowTransitionProperties`
**Summary:** Get workflow transition properties

This will be removed on [June 1, 2026](https://developer.atlassian.com/cloud/jira/platform/changelog/#CHANGE-2570); fetch transition properties from [Bulk get workflows](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflows-post) instead.

Returns the properties on a workflow transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/zIh

## Parameters
- `transitionId` [path] (required) integer(int64) — The ID of the transition. To get the ID, view the workflow in text mode in the Jira administration console. The ID is sh
- `includeReservedKeys` [query] boolean — Some properties with keys that have the *jira.* prefix are reserved, which means they are not editable. To include these
- `key` [query] string — The key of the property being returned, also known as the name of the property. If this parameter is not specified, all 
- `workflowName` [query] (required) string — The name of the workflow that the transition belongs to.
- `workflowMode` [query] string — The workflow status. Set to *live* for active and inactive workflows, or *draft* for draft workflows.

## Responses
- 200: object:
  - `id`: string
  - `key`: string
  - `value` (required): string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have admin permission
- 404: Returned if the workflow transition or property is not found.
