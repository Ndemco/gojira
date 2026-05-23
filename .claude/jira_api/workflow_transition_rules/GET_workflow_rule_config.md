# GET /rest/api/3/workflow/rule/config
**operationId:** `getWorkflowTransitionRuleConfigurations`
**Summary:** Get workflow transition rule configurations

Returns a [paginated](#pagination) list of workflows with transition rules. The workflows can be filtered to return only those containing workflow transition rules:

 *  of one or more transition rule types, such as [workflow post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/).
 *  matching one or more transition rule keys.

Only workflows containing transition rules created by the calling [Connect](https://developer.atlassian.com/cloud/jira/platf

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `types` [query] (required) []string — The types of the transition rules to return.
- `keys` [query] []string — The transition rule class keys, as defined in the Connect or the Forge app descriptor, of the transition rules to return
- `workflowNames` [query] []string — The list of workflow names to filter by.
- `withTags` [query] []string — The list of `tags` to filter by.
- `draft` [query] boolean — **Deprecated:** Whether draft or published workflows are returned. If not provided, both workflow types are returned. Th
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts `transition`, which, 

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []WorkflowTransitionRules
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: Returned if any transition rule type is not supported.
- 503: Returned if we encounter a problem while trying to access the required data.
