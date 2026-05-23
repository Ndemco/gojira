# PUT /rest/api/3/workflow/rule/config
**operationId:** `updateWorkflowTransitionRuleConfigurations`
**Summary:** Update workflow transition rule configurations

Updates configuration of workflow transition rules. The following rule types are supported:

 *  [post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/)
 *  [conditions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-condition/)
 *  [validators](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-validator/)

Only rules created by the calling [Connect](https://developer.atlassian.com/cloud/jira/platform/index/#c

## Request Body
Content-Type: `application/json`
object:
  - `workflows` (required): []WorkflowTransitionRules

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
- 503: Returned if we encounter a problem while trying to access the required data.
