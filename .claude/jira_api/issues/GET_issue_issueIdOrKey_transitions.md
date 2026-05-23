# GET /rest/api/3/issue/{issueIdOrKey}/transitions
**operationId:** `getTransitions`
**Summary:** Get transitions

Returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.

Note, if a request is made for a transition that does not exist or cannot be performed on the issue, given its status, the response will return any empty transitions list.

This operation can be accessed anonymously.

**[Permissions](#permissions) required: A list or transition is returned only when the user has:**

 *  *Browse projects* [project permission](https://conflue

## Parameters
- `issueIdOrKey` [path] (required) string — The ID or key of the issue.
- `expand` [query] string — Use [expand](#expansion) to include additional information about transitions in the response. This parameter accepts `tr
- `transitionId` [query] string — The ID of the transition.
- `skipRemoteOnlyCondition` [query] boolean — Whether transitions with the condition *Hide From User Condition* are included in the response. Available to Connect and
- `includeUnavailableTransitions` [query] boolean — Whether details of transitions that fail a condition are included in the response
- `sortByOpsBarAndStatus` [query] boolean — Whether the transitions are sorted by ops-bar sequence value first then category order (Todo, In Progress, Done) or only

## Responses
- 200: object:
  - `expand`: string
  - `transitions`: []IssueTransition
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the issue is not found or the user does not have permission to view it.
