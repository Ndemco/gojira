# POST /rest/api/3/issue
**operationId:** `createIssue`
**Summary:** Create issue

Creates an issue or, where the option to create subtasks is enabled in Jira, a subtask. A transition may be applied, to move the issue or subtask to a workflow step other than the default start step, and issue properties set.

The content of the issue or subtask is defined using `update` and `fields`. The fields that can be set in the issue or subtask are determined using the [ Get create issue metadata](#api-rest-api-3-issue-createmeta-get). These are the same fields that appear on the issue's 

## Parameters
- `updateHistory` [query] boolean — Whether the project in which the issue is created is added to the user's **Recently viewed** project list, as shown unde

## Request Body
Content-Type: `application/json`
object:
  - `fields`: object
  - `historyMetadata`: allOf(HistoryMetadata)
  - `properties`: []EntityProperty
  - `transition`: allOf(IssueTransition)
  - `update`: object

## Responses
- 201: object:
  - `id`: string
  - `key`: string
  - `self`: string
  - `transition`: allOf(NestedResponse)
  - `watchers`: allOf(NestedResponse)
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 422: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
