# POST /rest/api/3/issue/bulk
**operationId:** `createIssues`
**Summary:** Bulk create issue

Creates upto **50** issues and, where the option to create subtasks is enabled in Jira, subtasks. Transitions may be applied, to move the issues or subtasks to a workflow step other than the default start step, and issue properties set.

The content of each issue or subtask is defined using `update` and `fields`. The fields that can be set in the issue or subtask are determined using the [ Get create issue metadata](#api-rest-api-3-issue-createmeta-get). These are the same fields that appear on 

## Request Body
Content-Type: `application/json`
object:
  - `issueUpdates`: []IssueUpdateDetails

## Responses
- 201: object:
  - `errors`: []BulkOperationErrorResult
  - `issues`: []CreatedIssue
- 400: object:
  - `errors`: []BulkOperationErrorResult
  - `issues`: []CreatedIssue
- 401: Returned if the authentication credentials are incorrect or missing.
