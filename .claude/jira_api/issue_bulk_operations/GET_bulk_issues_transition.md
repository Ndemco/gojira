# GET /rest/api/3/bulk/issues/transition
**operationId:** `getAvailableTransitions`
**Summary:** Get available transitions

Use this API to retrieve a list of transitions available for the specified issues that can be used or bulk transition operations. You can submit either single or multiple issues in the query to obtain the available transitions.

The response will provide the available transitions for issues, organized by their respective workflows. **Only the transitions that are common among the issues within that workflow and do not involve any additional field updates will be included.** For bulk transitions 

## Parameters
- `issueIdsOrKeys` [query] (required) string — Comma (,) separated Ids or keys of the issues to get transitions available for them.
- `endingBefore` [query] string — (Optional)The end cursor for use in pagination.
- `startingAfter` [query] string — (Optional)The start cursor for use in pagination.

## Responses
- 200: object:
  - `availableTransitions`: []IssueBulkTransitionForWorkflow
  - `endingBefore`: string
  - `startingAfter`: string
- 400: object:
  - `errors`: []ErrorMessage
- 401: object:
  - `errors`: []ErrorMessage
- 403: object:
  - `errors`: []ErrorMessage
