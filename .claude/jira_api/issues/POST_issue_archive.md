# POST /rest/api/3/issue/archive
**operationId:** `archiveIssuesAsync`
**Summary:** Archive issue(s) by JQL

Enables admins to archive up to 100,000 issues in a single request using JQL, returning the URL to check the status of the submitted request.

You can use the [get task](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-tasks/#api-rest-api-3-task-taskid-get) and [cancel task](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-tasks/#api-rest-api-3-task-taskid-cancel-post) APIs to manage the request.

**Note that:**

 *  you can't archive subtasks directly, 

## Request Body
Content-Type: `application/json`
object:
  - `jql`: string

## Responses
- 202: string
- 400: any
- 401: any
- 403: any
- 412: any
