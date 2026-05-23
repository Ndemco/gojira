# PUT /rest/api/3/workflow/rule/config/delete
**operationId:** `deleteWorkflowTransitionRuleConfigurations`
**Summary:** Delete workflow transition rule configurations

Deletes workflow transition rules from one or more workflows. These rule types are supported:

 *  [post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/)
 *  [conditions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-condition/)
 *  [validators](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-validator/)

Only rules created by the calling Connect app can be deleted.

**Note:** The `draft` parameter in the 

## Request Body
Content-Type: `application/json`
object:
  - `workflows` (required): []WorkflowTransitionRulesDetails

## Responses
- 200: object:
  - `updateResults` (required): []WorkflowTransitionRulesUpdateErrorDetails
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
